package model

import (
	"github.com/_dev/exemplo-api-rest/model/entity"
)

// ListarAluno Retorna lista total alunos registrados.
func ListarAluno(db *DB) ([]*entity.Aluno, error) {
	rows, err := db.Query("SELECT * FROM aluno")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lista := make([]*entity.Aluno, 0)
	for rows.Next() {
		item := new(entity.Aluno)
		err := rows.Scan(&item.IDPessoa, &item.IDTurma)
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

// BuscarAluno Retorna um aluno especifico.
func BuscarAluno(db *DB, idPessoa int64) (*entity.Aluno, error) {
	row := db.QueryRow("SELECT * FROM aluno WHERE id_pessoa=$1", idPessoa)

	item := new(entity.Aluno)
	err := row.Scan(&item.IDPessoa, &item.IDTurma)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// InserirAluno Insere um novo registro aluno na base.
func InserirAluno(db *DB, entidade entity.Aluno) (int64, error) {
	sqlStatement := "INSERT INTO aluno (id_pessoa, id_turma) VALUES ($1, $2) RETURNING id_pessoa"

	var idretorno int64
	err := db.QueryRow(sqlStatement, entidade.IDPessoa, entidade.IDTurma).Scan(&idretorno)
	if err != nil {
		return 0, err
	}

	return idretorno, nil
}

// AtualizarAluno Atualiza um registro aluno da base.
func AtualizarAluno(db *DB, entidade entity.Aluno, entidadepessoa entity.Pessoa) error {
	sqlStatement := "UPDATE aluno SET id_turma=$2 WHERE id_pessoa=$1"
	_, err := db.Exec(sqlStatement, entidade.IDPessoa, entidade.IDTurma)
	if err != nil {
		return err
	}
	if err = AtualizarPessoa(db, entidadepessoa); err != nil {
		return err
	}
	return nil
}
