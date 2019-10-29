package entity

import (
	"encoding/json"
)

// Student Entity.
type Student struct {
	Person Person `json:"person,omitempty"`
	Class  Class `json:"class,omitempty"`
}

// New - Loads a new Student structure.
func (entity *Student) New(person Person, class Class) {
	*entity = Student{person, class}
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
		"Person": entity.Person.ToString(),
		"Class":  entity.Class.ToString(),
	}
	retorno := ToString("Student", campos)
	return retorno
}
