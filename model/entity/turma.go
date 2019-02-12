package entity

import (
	"encoding/json"
	"rest-api/util"
	"strconv"
	"time"
)

// Turma Entidade.
type Turma struct {
	ID           int64     `json:"id,omitempty"`
	IDCurso      int64     `json:"idcurso,omitempty"`
	DataInicio   time.Time `json:"datainicio,omitempty"`
	DataFim      time.Time `json:"datafim,omitempty"`
	DataCadastro time.Time `json:"datacadastro,omitempty"`
}

// New Carrega uma nova estrutura Turma.
func (entidade *Turma) New(id int64, idCurso int64, dataInicio time.Time, dataFim time.Time, dataCadastro time.Time) {
	*entidade = Turma{id, idCurso, dataInicio, dataFim, dataCadastro}
}

// Decoder Decodifica JSON para estrutura.
func (entidade *Turma) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entidade); err != nil {
		return err
	}
	return nil
}

// ToString Retorna string com as informacoes de Turma.
func (entidade *Turma) ToString() string {
	campos := map[string]string{
		"ID":           strconv.FormatInt(entidade.ID, 10),
		"IDCurso":      strconv.FormatInt(entidade.IDCurso, 10),
		"DataInicio":   util.FormatarDataHora(entidade.DataInicio),
		"DataFim":      util.FormatarDataHora(entidade.DataFim),
		"DataCadastro": util.FormatarDataHora(entidade.DataCadastro),
	}
	retorno := ToString("Turma", campos)
	return retorno
}
