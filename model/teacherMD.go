package model

import (
	"github.com/_dev/exemplo-api-rest/model/entity"
)

// ListarProfessor Retorna lista total professores registrados.
func ListarProfessor(db *DB) ([]*entity.Professor, error) {
	rows, err := db.Query("SELECT * FROM professor")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lista := make([]*entity.Professor, 0)
	for rows.Next() {
		item := new(entity.Professor)
		err := rows.Scan(&item.IDPessoa, &item.IDCurso)
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

// BuscarProfessor Retorna um professor especifico.
func BuscarProfessor(db *DB, id int64) (*entity.Professor, error) {
	row := db.QueryRow("SELECT * FROM professor WHERE id=$1", id)

	item := new(entity.Professor)
	err := row.Scan(&item.IDPessoa, &item.IDCurso)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// InserirProfessor Insere um novo registro professor na base.
func InserirProfessor(db *DB, entidade entity.Professor) (int64, error) {
	sqlStatement := "INSERT INTO professor (id_pessoa, id_curso) VALUES ($1, $2) RETURNING id_pessoa"

	var idretorno int64
	err := db.QueryRow(sqlStatement, entidade.IDPessoa, entidade.IDCurso).Scan(&idretorno)
	if err != nil {
		return 0, err
	}

	return idretorno, nil
}

// AtualizarProfessor Atualiza um registro professor da base.
func AtualizarProfessor(db *DB, entidade entity.Professor, entidadepessoa entity.Pessoa) error {
	sqlStatement := "UPDATE professor SET id_curso=$2 WHERE id_pessoa=$1"
	_, err := db.Exec(sqlStatement, entidade.IDPessoa, entidade.IDCurso)
	if err != nil {
		return err
	}
	if err = AtualizarPessoa(db, entidadepessoa); err != nil {
		return err
	}
	return nil
}
