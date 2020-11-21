package model

import (
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/util"

)

// Course ...
type Course struct {
	ID               int       `db:"id" json:"id,omitempty"`
	Name             string    `db:"name" json:"name,omitempty"`
	Description      string    `db:"description" json:"description,omitempty"`
	RegistrationDate time.Time `db:"registration_date" json:"registration_date,omitempty"`
}

// ToString ...
func (entity *Course) ToString() string {
	result, _ := getJSONSerilizer(entity)
	return result
}

// GetTableName ...
func (entity *Course) GetTableName() string {
	return util.GetType(entity)
}
