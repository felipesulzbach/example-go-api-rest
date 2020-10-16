package model

import (
	"encoding/json"

)

// Teacher Entity.
type Teacher struct {
	Person Person `json:"person,omitempty"`
	Course Course `json:"course,omitempty"`
}

// New - Loads a new Teacher structure.
func (entity *Teacher) New(person Person, course Course) {
	*entity = Teacher{person, course}
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
		"Person": entity.Person.ToString(),
		"Course": entity.Course.ToString(),
	}
	retorno := ToString("Teacher", campos)
	return retorno
}
