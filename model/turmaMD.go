package model

import (
	"github.com/_dev/exemplo-api-rest/model/entity"
)

// NextIDTurma Retorna o proximo ID.
func NextIDTurma(db *DB) (int64, error) {
	row := db.QueryRow("SELECT (MAX(id) + 1) FROM turma")

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// ListarTurma Retorna lista total turmas registrados.
func ListarTurma(db *DB) ([]*entity.Turma, error) {
	rows, err := db.Query("SELECT * FROM turma")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lista := make([]*entity.Turma, 0)
	for rows.Next() {
		item := new(entity.Turma)
		err := rows.Scan(&item.ID, &item.IDCurso, &item.DataInicio, &item.DataFim, &item.DataCadastro)
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

// BuscarTurma Retorna um curso especifico.
func BuscarTurma(db *DB, id int64) (*entity.Turma, error) {
	row := db.QueryRow("SELECT * FROM turma WHERE id=$1", id)

	item := new(entity.Turma)
	err := row.Scan(&item.ID, &item.IDCurso, &item.DataInicio, &item.DataFim, &item.DataCadastro)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// InserirTurma Insere um novo registro turma na base.
func InserirTurma(db *DB, entidade entity.Turma) (int64, error) {
	sqlStatement := "INSERT INTO turma (id, id_curso, data_inicio, data_fim, data_cadastro) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var idretorno int64
	err := db.QueryRow(sqlStatement, entidade.ID, entidade.IDCurso, entidade.DataInicio, entidade.DataFim, entidade.DataCadastro).Scan(&idretorno)
	if err != nil {
		return 0, err
	}

	return idretorno, nil
}

// AtualizarTurma Atualiza um registro turma da base.
func AtualizarTurma(db *DB, entidade entity.Turma) error {
	sqlStatement := "UPDATE turma SET id_curso=$2, data_inicio=$3, data_fim=$4, data_cadastro=$5 WHERE id=$1"
	_, err := db.Exec(sqlStatement, entidade.ID, entidade.IDCurso, entidade.DataInicio, entidade.DataFim, entidade.DataCadastro)
	if err != nil {
		return err
	}
	return nil
}
