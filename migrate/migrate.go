package migrate

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/migrate/repository"
	"github.com/felipesulzbach/exemplo-files-process/util"

)

const (
	filePath      = "db/migration"
	fileExtention = "sql"
)

type file struct {
	fileName string
	filePath string
}

// LoadDatabaseStructure ...
func LoadDatabaseStructure() {
	log.Println("Starting migration...")

	db, err := repository.OpenDB()
	if err != nil {
		log.Panic(err)
	}

	sqlFiles := getFilesPath(filePath, fileExtention)

	if err = executeQuerySchema("CREATE SCHEMA IF NOT EXISTS fs_auto", db); err != nil {
		log.Panic(err)
	}
	if err = executeQuerySchema("CREATE TABLE IF NOT EXISTS fs_auto.migrate (migrate_name CHARACTER VARYING(255),registration_date TIMESTAMP)", db); err != nil {
		log.Panic(err)
	}

	for _, item := range sqlFiles {
		isExecuted, err := isFileExecuted(item.fileName, db)
		if err != nil {
			log.Panic(err)
		}
		if isExecuted {
			continue
		}

		if err = executeQuerysInFile(item.filePath, db); err != nil {
			log.Panic(err)
		}

		if err = insertMigrate(item.fileName, db); err != nil {
			log.Panic(err)
		}

		log.Printf(">>> %s.\n", item.filePath)
	}

	db.CloseDB()

	log.Println("Migration completed.")
}

func getFilesPath(filePath, fileExtention string) []file {
	var response []file
	var filesPath []file

	err := filepath.Walk(filePath, visit(&filesPath))
	if err != nil {
		panic(err)
	}
	for _, item := range filesPath {
		extension, _ := util.GetFileExtension(item.filePath)
		if extension == fileExtention {
			response = append(response, file{item.fileName, strings.ReplaceAll(item.filePath, "\\", "\\\\")})
		}
	}

	return response
}

func visit(filesPath *[]file) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		*filesPath = append(*filesPath, file{info.Name(), path})

		return nil
	}
}

func executeQuerysInFile(sqlFile string, db *repository.DB) error {
	file, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		tx.Rollback()
	}()

	for _, q := range strings.Split(string(file), ";") {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}
		if _, err := tx.Exec(q); err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

func executeQuerySchema(query string, db *repository.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		tx.Rollback()
	}()

	if _, err := tx.Exec(query); err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func isFileExecuted(fileName string, db *repository.DB) (bool, error) {
	row := db.QueryRow("SELECT COUNT(migrate_name) FROM fs_auto.migrate WHERE migrate_name=$1", fileName)
	item := 0
	err := row.Scan(&item)
	if err != nil {
		return false, err
	}
	return item == 1, nil
}

func insertMigrate(fileName string, db *repository.DB) error {
	sqlStatement := "INSERT INTO fs_auto.migrate (migrate_name,registration_date) VALUES ($1, $2)"
	_, err := db.Query(sqlStatement, fileName, time.Now())
	if err != nil {
		return err
	}

	return nil
}
