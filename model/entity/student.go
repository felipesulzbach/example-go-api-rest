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
func (entity *Student) New(personID int64, classID int64) {
	*entity = Student{personID, classID}
}

// Decoder - Decodes JSON for structure.
func (entity *Student) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entity); err != nil {
		return err
	}
	return nil
}

// ToString - Returns string with Student information.
func (entity *Student) ToString() string {
	campos := map[string]string{
		"PersonID": strconv.FormatInt(entity.PersonID, 10),
		"ClassID":  strconv.FormatInt(entity.ClassID, 10),
	}
	retorno := ToString("Student", campos)
	return retorno
}
