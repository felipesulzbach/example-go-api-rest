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
func (entity *Class) New(id int64, courseID int64, startDate time.Time, endDate time.Time, registrationDate time.Time) {
	*entity = Class{id, courseID, startDate, endDate, registrationDate}
}

// Decoder - Decodes JSON for structure.
func (entity *Class) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entity); err != nil {
		return err
	}
	return nil
}

// ToString - Returns string with Person information.
func (entity *Class) ToString() string {
	campos := map[string]string{
		"ID":               strconv.FormatInt(entity.ID, 10),
		"CourseID":         strconv.FormatInt(entity.CourseID, 10),
		"StartDate":        util.FormatDateTime(entity.StartDate),
		"EndDate":          util.FormatDateTime(entity.EndDate),
		"RegistrationDate": util.FormatDateTime(entity.RegistrationDate),
	}
	retorno := ToString("Class", campos)
	return retorno
}
