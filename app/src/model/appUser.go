package model

import (
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/util"

)

// AppUser Entity.
type AppUser struct {
	ID               int64     `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Password         string    `json:"password,omitempty"`
	Profile          Profile   `json:"profile,omitempty"`
	Person           Person    `json:"person,omitempty"`
	StartDate        time.Time `json:"start_date,omitempty"`
	EndDate          time.Time `json:"end_date,omitempty"`
	RegistrationDate time.Time `json:"registration_date,omitempty"`
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
