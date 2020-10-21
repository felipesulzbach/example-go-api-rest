package model

import (
	"time"

)

// Person Entity.
type Person struct {
	entityModel
	ID               int64     `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Cpf              string    `json:"cpf,omitempty"`
	CellPhone        string    `json:"cell_phone,omitempty"`
	City             string    `json:"city,omitempty"`
	ZipCode          string    `json:"zip_code,omitempty"`
	Address          string    `json:"address,omitempty"`
	RegistrationDate time.Time `json:"registration_date,omitempty"`
}
