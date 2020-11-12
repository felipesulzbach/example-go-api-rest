package model

import (
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/util"

)

// Course ...
type Course struct {
	ID               int     `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Description      string    `json:"description,omitempty"`
	RegistrationDate time.Time `json:"registration_date,omitempty"`
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
