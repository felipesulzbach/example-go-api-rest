package model

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/util"

)

// Course Entity.
type Course struct {
	ID               int64     `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Description      string    `json:"description,omitempty"`
	RegistrationDate time.Time `json:"registrationDate,omitempty"`
}

// New - Loads a new Course structure.
func (entity *Course) New(id int64, name string, description string, registrationDate time.Time) {
	*entity = Course{id, name, description, registrationDate}
}

// Decoder - Decodes JSON for structure.
func (entity *Course) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entity); err != nil {
		return err
	}
	return nil
}

// ToString - Returns string with Course information.
func (entity *Course) ToString() string {
	campos := map[string]string{
		"ID":               strconv.FormatInt(entity.ID, 10),
		"Name":             entity.Name,
		"Description":      entity.Description,
		"RegistrationDate": util.FormatDateTime(entity.RegistrationDate),
	}
	retorno := ToString("Course", campos)
	return retorno
}
