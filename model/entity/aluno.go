package entity

import (
	"encoding/json"
	"strconv"
)

// Aluno Entidade.
type Aluno struct {
	IDPessoa int64 `json:"idpessoa,omitempty"`
	IDTurma  int64 `json:"idturma,omitempty"`
}

// New Carrega uma nova estrutura Aluno.
func (entidade *Aluno) New(idPessoa int64, idTurma int64) {
	*entidade = Aluno{idPessoa, idTurma}
}

// Decoder Decodifica JSON para estrutura.
func (entidade *Aluno) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entidade); err != nil {
		return err
	}
	return nil
}

// ToString Retorna string com as informacoes de Aluno.
func (entidade *Aluno) ToString() string {
	campos := map[string]string{
		"IDPessoa": strconv.FormatInt(entidade.IDPessoa, 10),
		"IDTurma":  strconv.FormatInt(entidade.IDTurma, 10),
	}
	retorno := ToString("Aluno", campos)
	return retorno
}
