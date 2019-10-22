package entity

import (
	"encoding/json"
	"strconv"
)

// Student Entity.
type Student struct {
	PersonID int64 `json:"personID,omitempty"`
	ClassID  int64 `json:"classID,omitempty"`
}

// New - Loads a new Student structure.
func (entidade *Student) New(personID int64, classID int64) {
	*entidade = Student{personID, classID}
}

// Decoder - Decodes JSON for structure.
func (entidade *Student) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entidade); err != nil {
		return err
	}
	return nil
}

// ToString - Returns string with Student information.
func (entidade *Student) ToString() string {
	campos := map[string]string{
		"PersonID": strconv.FormatInt(entidade.PersonID, 10),
		"ClassID":  strconv.FormatInt(entidade.ClassID, 10),
	}
	retorno := ToString("Student", campos)
	return retorno
}
