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
func (entity *Teacher) New(personID int64, courseID int64) {
	*entity = Teacher{personID, courseID}
}

// Decoder - Decodes JSON for structure.
func (entity *Teacher) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entity); err != nil {
		return err
	}
	return nil
}

// ToString - Returns string with Teacher information.
func (entity *Teacher) ToString() string {
	campos := map[string]string{
		"Pessoa":   strconv.FormatInt(entity.PersonID, 10),
		"CourseID": strconv.FormatInt(entity.CourseID, 10),
	}
	retorno := ToString("Teacher", campos)
	return retorno
}
