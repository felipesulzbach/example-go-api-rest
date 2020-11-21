package model

import "github.com/felipesulzbach/exemplo-api-rest/app/src/util"

// Teacher Entity.
type Teacher struct {
	Person Person `db:"person" json:"person,omitempty"`
	Course Course `db:"course" json:"course,omitempty"`
}

// ToString ...
func (entity *Teacher) ToString() string {
	result, _ := getJSONSerilizer(entity)
	return result
}

// GetTableName ...
func (entity *Teacher) GetTableName() string {
	return util.GetType(entity)
}
