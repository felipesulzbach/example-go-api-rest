package entity

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/_dev/exemplo-api-rest/util"
)

// Class Entity.
type Class struct {
	ID               int64     `json:"id,omitempty"`
	CourseID         int64     `json:"courseID,omitempty"`
	StartDate        time.Time `json:"startDate,omitempty"`
	EndDate          time.Time `json:"endDate,omitempty"`
	RegistrationDate time.Time `json:"registrationDate,omitempty"`
}

// New - Loads a new Person structure.
func (entidade *Class) New(id int64, courseID int64, startDate time.Time, endDate time.Time, registrationDate time.Time) {
	*entidade = Class{id, courseID, startDate, endDate, registrationDate}
}

// Decoder - Decodes JSON for structure.
func (entidade *Class) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entidade); err != nil {
		return err
	}
	return nil
}

// ToString - Returns string with Person information.
func (entidade *Class) ToString() string {
	campos := map[string]string{
		"ID":               strconv.FormatInt(entidade.ID, 10),
		"CourseID":         strconv.FormatInt(entidade.CourseID, 10),
		"StartDate":        util.FormatDateTime(entidade.StartDate),
		"EndDate":          util.FormatDateTime(entidade.EndDate),
		"RegistrationDate": util.FormatDateTime(entidade.RegistrationDate),
	}
	retorno := ToString("Class", campos)
	return retorno
}
