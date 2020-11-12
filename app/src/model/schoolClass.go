package model

import (
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/util"

)

// SchoolClass ...
type SchoolClass struct {
	ID               int64     `json:"id,omitempty"`
	Course           Course    `json:"course,omitempty"`
	StartDate        time.Time `json:"start_date,omitempty"`
	EndDate          time.Time `json:"end_date,omitempty"`
	RegistrationDate time.Time `json:"registration_date,omitempty"`
}

// ToString ...
func (entity *SchoolClass) ToString() string {
	result, _ := getJSONSerilizer(entity)
	return result
}

// GetTableName ...
func (entity *SchoolClass) GetTableName() string {
	return util.GetType(entity)
}
