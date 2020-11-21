package model

import (
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/util"

)

// AppUser Entity.
type AppUser struct {
	ID               int64     `db:"id" json:"id,omitempty"`
	Name             string    `db:"name" json:"name,omitempty"`
	Password         string    `db:"password" json:"password,omitempty"`
	Profile          Profile   `db:"profile" json:"profile,omitempty"`
	Person           Person    `db:"person" json:"person,omitempty"`
	StartDate        time.Time `db:"start_date" json:"start_date,omitempty"`
	EndDate          time.Time `db:"end_date" json:"end_date,omitempty"`
	RegistrationDate time.Time `db:"registration_date" json:"registration_date,omitempty"`
}

// ToString ...
func (entity *AppUser) ToString() (string, error) {
	result, err := getJSONSerilizer(entity)
	if err != nil {
		return "", err
	}
	return result, nil
}

// GetTableName ...
func (entity *AppUser) GetTableName() string {
	return util.GetType(entity)
}
