package model

import (
	"time"

	"github.com/felipesulzbach/example-go-api-rest/src/domain/util"

)

// Person Entity.
type Person struct {
	ID               int64     `db:"id" json:"id,omitempty"`
	Name             string    `db:"name" json:"name,omitempty"`
	Cpf              string    `db:"cpf" json:"cpf,omitempty"`
	CellPhone        string    `db:"cell_phone" json:"cell_phone,omitempty"`
	City             string    `db:"city" json:"city,omitempty"`
	ZipCode          string    `db:"zip_code" json:"zip_code,omitempty"`
	Address          string    `db:"address" json:"address,omitempty"`
	RegistrationDate time.Time `db:"registration_date" json:"registration_date,omitempty"`
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
