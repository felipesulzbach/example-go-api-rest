package model

import "time"

// AppUser Entity.
type AppUser struct {
	entityModel
	ID               int64     `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Password         string    `json:"password,omitempty"`
	Profile          Profile   `json:"profile,omitempty"`
	Person           Person    `json:"person,omitempty"`
	StartDate        time.Time `json:"start_date,omitempty"`
	EndDate          time.Time `json:"end_date,omitempty"`
	RegistrationDate time.Time `json:"registration_date,omitempty"`
}
