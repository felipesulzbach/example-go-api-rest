package model

import (
	"time"

)

// Course Entity.
type Course struct {
	entityModel
	ID               int64     `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Description      string    `json:"description,omitempty"`
	RegistrationDate time.Time `json:"registration_date,omitempty"`
}
