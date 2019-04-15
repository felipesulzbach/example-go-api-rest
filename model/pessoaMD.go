package model

import (
	"exemplo-api-rest/model/entity"
)

// NextIDPessoa Retorna o proximo ID.
func NextIDPessoa(db *DB) (int64, error) {
	row := db.QueryRow("SELECT (MAX(id) + 1) FROM pessoa")

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// ListarPessoa Retorna lista total pessoas registradas.
func ListarPessoa(db *DB) ([]*entity.Pessoa, error) {
	rows, err := db.Query("SELECT * FROM pessoa")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lista := make([]*entity.Pessoa, 0)
	for rows.Next() {
		item := new(entity.Pessoa)
		err := rows.Scan(&item.ID, &item.Nome, &item.NumeroCpf, &item.NumeroCelular, &item.Cidade, &item.NumeroCep, &item.Endereco, &item.DataCadastro)
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

// BuscarPessoa Retorna um pessoa especifico.
func BuscarPessoa(db *DB, id int64) (*entity.Pessoa, error) {
	row := db.QueryRow("SELECT * FROM pessoa WHERE id=$1", id)

	item := new(entity.Pessoa)
	err := row.Scan(&item.ID, &item.Nome, &item.NumeroCpf, &item.NumeroCelular, &item.Cidade, &item.NumeroCep, &item.Endereco, &item.DataCadastro)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// InserirPessoa Insere um novo registro pessoa na base.
func InserirPessoa(db *DB, entidade entity.Pessoa) (int64, error) {
	sqlStatement := "INSERT INTO pessoa (id, nome, numero_cpf, numero_celular, cidade, numero_cep, endereco, data_cadastro) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	var idretorno int64
	err := db.QueryRow(sqlStatement, entidade.ID, entidade.Nome, entidade.NumeroCpf, entidade.NumeroCelular, entidade.Cidade, entidade.NumeroCep, entidade.Endereco, entidade.DataCadastro).Scan(&idretorno)
	if err != nil {
		return 0, err
	}

	return idretorno, nil
}

// AtualizarPessoa Atualiza um registro pessoa da base.
func AtualizarPessoa(db *DB, entidade entity.Pessoa) error {
	sqlStatement := "UPDATE pessoa SET nome=$2, numero_cpf=$3, numero_celular=$4, cidade=$5, numero_cep=$6, endereco=$7, data_cadastro=$8 WHERE id=$1"
	_, err := db.Exec(sqlStatement, entidade.ID, entidade.Nome, entidade.NumeroCpf, entidade.NumeroCelular, entidade.Cidade, entidade.NumeroCep, entidade.Endereco, entidade.DataCadastro)
	if err != nil {
		return err
	}
	return nil
}
