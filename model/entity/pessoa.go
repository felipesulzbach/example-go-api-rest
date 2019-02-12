package entity

import (
	"encoding/json"
	"rest-api/util"
	"strconv"
	"time"
)

// Pessoa Entidade.
type Pessoa struct {
	ID            int64     `json:"id,omitempty"`
	Nome          string    `json:"nome,omitempty"`
	NumeroCpf     string    `json:"numerocpf,omitempty"`
	NumeroCelular string    `json:"numerocelular,omitempty"`
	Cidade        string    `json:"cidade,omitempty"`
	NumeroCep     string    `json:"numerocep,omitempty"`
	Endereco      string    `json:"endereco,omitempty"`
	DataCadastro  time.Time `json:"datacadastro,omitempty"`
}

// New Carrega uma nova estrutura Pessoa.
func (entidade *Pessoa) New(id int64, nome string, numeroCpf string, numeroCelular string, cidade string, numeroCep string, endereco string, dataCadastro time.Time) {
	*entidade = Pessoa{id, nome, numeroCpf, numeroCelular, cidade, numeroCep, endereco, dataCadastro}
}

// Decoder Decodifica JSON para estrutura.
func (entidade *Pessoa) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entidade); err != nil {
		return err
	}
	return nil
}

// ToString Retorna string com as informacoes de Pessoa.
func (entidade *Pessoa) ToString() string {
	campos := map[string]string{
		"ID":            strconv.FormatInt(entidade.ID, 10),
		"Nome":          entidade.Nome,
		"NumeroCpf":     entidade.NumeroCpf,
		"NumeroCelular": entidade.NumeroCelular,
		"Cidade":        entidade.Cidade,
		"NumeroCep":     entidade.NumeroCep,
		"Endereco":      entidade.Endereco,
		"DataCadastro":  util.FormatarDataHora(entidade.DataCadastro),
	}
	retorno := ToString("Pessoa", campos)
	return retorno
}
