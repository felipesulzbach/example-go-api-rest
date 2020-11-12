package model

import (
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/util"

)

// Person Entity.
type Person struct {
	ID               int64     `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Cpf              string    `json:"cpf,omitempty"`
	CellPhone        string    `json:"cell_phone,omitempty"`
	City             string    `json:"city,omitempty"`
	ZipCode          string    `json:"zip_code,omitempty"`
	Address          string    `json:"address,omitempty"`
	RegistrationDate time.Time `json:"registration_date,omitempty"`
}

// ToString ...
func (entity *Person) ToString() string {
	result, _ := getJSONSerilizer(entity)
	return result
}

// GetTableName ...
func (entity *Person) GetTableName() string {
	return util.GetType(entity)
}
