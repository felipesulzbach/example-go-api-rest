package entity

import (
	"encoding/json"
	"exemplo-api-rest/util"
	"strconv"
	"time"
)

// Curso Entidade.
type Curso struct {
	ID           int64     `json:"id,omitempty"`
	Nome         string    `json:"nome,omitempty"`
	Descricao    string    `json:"descricao,omitempty"`
	DataCadastro time.Time `json:"datacadastro,omitempty"`
}

// New Carrega uma nova estrutura Curso.
func (entidade *Curso) New(id int64, nome string, descricao string, dataCadastro time.Time) {
	*entidade = Curso{id, nome, descricao, dataCadastro}
}

// Decoder Decodifica JSON para estrutura.
func (entidade *Curso) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entidade); err != nil {
		return err
	}
	return nil
}

// ToString Retorna string com as informacoes de Curso.
func (entidade *Curso) ToString() string {
	campos := map[string]string{
		"ID":           strconv.FormatInt(entidade.ID, 10),
		"Nome":         entidade.Nome,
		"Descricao":    entidade.Descricao,
		"DataCadastro": util.FormatarDataHora(entidade.DataCadastro),
	}
	retorno := ToString("Curso", campos)
	return retorno
}
