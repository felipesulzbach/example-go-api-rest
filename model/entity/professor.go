package entity

import (
	"encoding/json"
	"strconv"
)

// Professor Entidade.
type Professor struct {
	IDPessoa int64 `json:"idpessoa,omitempty"`
	IDCurso  int64 `json:"idcurso,omitempty"`
}

// New Carrega uma nova estrutura Professor.
func (entidade *Professor) New(idPessoa int64, idCurso int64) {
	*entidade = Professor{idPessoa, idCurso}
}

// Decoder Decodifica JSON para estrutura.
func (entidade *Professor) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entidade); err != nil {
		return err
	}
	return nil
}

// ToString Retorna string com as informacoes de Professor.
func (entidade *Professor) ToString() string {
	campos := map[string]string{
		"Pessoa":  strconv.FormatInt(entidade.IDPessoa, 10),
		"IDCurso": strconv.FormatInt(entidade.IDCurso, 10),
	}
	retorno := ToString("Professor", campos)
	return retorno
}
