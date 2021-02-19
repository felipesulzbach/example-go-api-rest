package migrate

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/felipesulzbach/example-go-api-rest/src/infra/db"
	"github.com/felipesulzbach/exemplo-files-process/util"

)

const (
	filePath      = "postgres/migration"
	fileExtention = "sql"
)

// LoadDatabaseStructure ...
func LoadDatabaseStructure() {
	log.Println("MIGRATION starting...")

	postgres, err := db.Connect()
	if err != nil {
		log.Panic(err)
	}

	sqlFiles := _getFilesPath(filePath, fileExtention)

	if err = _executeQuerySchema("CREATE SCHEMA IF NOT EXISTS fs_auto", postgres); err != nil {
		log.Panic(err)
	}
	if err = _executeQuerySchema("CREATE TABLE IF NOT EXISTS fs_auto.migrate (migrate_name CHARACTER VARYING(255),registration_date TIMESTAMP)", postgres); err != nil {
		log.Panic(err)
	}

	for _, item := range sqlFiles {
		isExecuted, err := _isFileExecuted(item.fileName, postgres)
		if err != nil {
			log.Panic(err)
		}
		if isExecuted {
			continue
		}

		if err = _executeQuerysInFile(item.filePath, postgres); err != nil {
			log.Panic(err)
		}

		if err = _insertMigrate(item.fileName, postgres); err != nil {
			log.Panic(err)
		}

		log.Printf("Migrate %s.\n", item.fileName)
	}

	postgres.Disconnect()

	log.Println("MIGRATION Completed!")
}

func _getFilesPath(filePath, fileExtention string) []file {
	var response []file
	var filesPath []file

	err := filepath.Walk(filePath, _visit(&filesPath))
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

func _visit(filesPath *[]file) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		*filesPath = append(*filesPath, file{info.Name(), path})

		return nil
	}
}

func _executeQuerysInFile(sqlFile string, postgres *db.DB) error {
	file, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		return err
	}

	tx, err := postgres.Begin()
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

func _executeQuerySchema(query string, postgres *db.DB) error {
	tx, err := postgres.Begin()
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

func _isFileExecuted(fileName string, postgres *db.DB) (bool, error) {
	row := postgres.QueryRow("SELECT COUNT(migrate_name) FROM fs_auto.migrate WHERE migrate_name=$1", fileName)
	item := 0
	err := row.Scan(&item)
	if err != nil {
		return false, err
	}
	return item == 1, nil
}

func _insertMigrate(fileName string, postgres *db.DB) error {
	sqlStatement := "INSERT INTO fs_auto.migrate (migrate_name,registration_date) VALUES ($1, $2)"
	_, err := postgres.Query(sqlStatement, fileName, time.Now())
	if err != nil {
		return err
	}

	return nil
}
