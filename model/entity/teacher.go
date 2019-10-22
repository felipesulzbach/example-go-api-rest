package entity

import (
	"encoding/json"
	"strconv"
)

// Teacher Entity.
type Teacher struct {
	PersonID int64 `json:"personID,omitempty"`
	CourseID int64 `json:"courseID,omitempty"`
}

// New - Loads a new Teacher structure.
func (entidade *Teacher) New(personID int64, courseID int64) {
	*entidade = Teacher{personID, courseID}
}

// Decoder - Decodes JSON for structure.
func (entidade *Teacher) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entidade); err != nil {
		return err
	}
	return nil
}

// ToString - Returns string with Teacher information.
func (entidade *Teacher) ToString() string {
	campos := map[string]string{
		"Pessoa":   strconv.FormatInt(entidade.PersonID, 10),
		"CourseID": strconv.FormatInt(entidade.CourseID, 10),
	}
	retorno := ToString("Teacher", campos)
	return retorno
}
