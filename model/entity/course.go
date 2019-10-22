package entity

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/_dev/exemplo-api-rest/util"
)

// Course Entity.
type Course struct {
	ID               int64     `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Description      string    `json:"description,omitempty"`
	RegistrationDate time.Time `json:"registrationDate,omitempty"`
}

// New - Loads a new Course structure.
func (entidade *Course) New(id int64, name string, description string, registrationDate time.Time) {
	*entidade = Course{id, name, description, registrationDate}
}

// Decoder - Decodes JSON for structure.
func (entidade *Course) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entidade); err != nil {
		return err
	}
	return nil
}

// ToString - Returns string with Course information.
func (entidade *Course) ToString() string {
	campos := map[string]string{
		"ID":               strconv.FormatInt(entidade.ID, 10),
		"Name":             entidade.Name,
		"Description":      entidade.Description,
		"RegistrationDate": util.FormatDateTime(entidade.RegistrationDate),
	}
	retorno := ToString("Course", campos)
	return retorno
}
