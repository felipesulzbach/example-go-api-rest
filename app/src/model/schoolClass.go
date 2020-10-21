package model

import (
	"time"

)

// SchoolClass ...
type SchoolClass struct {
	entityModel
	ID               int64     `json:"id,omitempty"`
	Course           Course    `json:"course,omitempty"`
	StartDate        time.Time `json:"start_date,omitempty"`
	EndDate          time.Time `json:"end_date,omitempty"`
	RegistrationDate time.Time `json:"registration_date,omitempty"`
}
