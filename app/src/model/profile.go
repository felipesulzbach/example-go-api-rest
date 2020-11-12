package model

import "github.com/felipesulzbach/exemplo-api-rest/app/src/util"

// Profile Entity.
type Profile struct {
	ID               int64  `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Code             string `json:"code,omitempty"`
	AccessPermission string `json:"access_permission,omitempty"`
}

// ToString ...
func (entity *Profile) ToString() string {
	result, _ := getJSONSerilizer(entity)
	return result
}

// GetTableName ...
func (entity *Profile) GetTableName() string {
	return util.GetType(entity)
}
