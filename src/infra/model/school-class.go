package model

import (
	"time"

	"github.com/felipesulzbach/example-go-api-rest/src/domain/util"

)

// SchoolClass ...
type SchoolClass struct {
	ID               int64     `db:"id" json:"id,omitempty"`
	Course           Course    `db:"course" json:"course,omitempty"`
	StartDate        time.Time `db:"start_date" json:"start_date,omitempty"`
	EndDate          time.Time `db:"end_date" json:"end_date,omitempty"`
	RegistrationDate time.Time `db:"registration_date" json:"registration_date,omitempty"`
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
