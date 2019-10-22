package model

import (
	"github.com/_dev/exemplo-api-rest/model/entity"
)

// NextIDCurso Retorna o proximo ID.
func NextIDCurso(db *DB) (int64, error) {
	row := db.QueryRow("SELECT (MAX(id) + 1) FROM curso")

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// ListarCurso Retorna lista total cursos registrados.
func ListarCurso(db *DB) ([]*entity.Curso, error) {
	rows, err := db.Query("SELECT * FROM curso")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lista := make([]*entity.Curso, 0)
	for rows.Next() {
		item := new(entity.Curso)
		err := rows.Scan(&item.ID, &item.Nome, &item.Descricao, &item.DataCadastro)
		if err != nil {
			return nil, err
		}
		lista = append(lista, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return lista, nil
}

// BuscarCurso Retorna um curso especifico.
func BuscarCurso(db *DB, id int64) (*entity.Curso, error) {
	row := db.QueryRow("SELECT * FROM curso WHERE id=$1", id)

	item := new(entity.Curso)
	err := row.Scan(&item.ID, &item.Nome, &item.Descricao, &item.DataCadastro)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// InserirCurso Insere um novo registro curso na base.
func InserirCurso(db *DB, entidade entity.Curso) (int64, error) {
	sqlStatement := "INSERT INTO curso (id, nome, descricao, data_cadastro) VALUES ($1, $2, $3, $4) RETURNING id"
	var idretorno int64
	err := db.QueryRow(sqlStatement, entidade.ID, entidade.Nome, entidade.Descricao, entidade.DataCadastro).Scan(&idretorno)
	if err != nil {
		return 0, err
	}

	return idretorno, nil
}

// AtualizarCurso Atualiza um registro curso da base.
func AtualizarCurso(db *DB, entidade entity.Curso) error {
	sqlStatement := "UPDATE curso SET nome=$2, descricao=$3, data_cadastro=$4 WHERE id=$1"
	_, err := db.Exec(sqlStatement, entidade.ID, entidade.Nome, entidade.Descricao, entidade.DataCadastro)
	if err != nil {
		return err
	}
	return nil
}
