package model

import "github.com/felipesulzbach/example-go-api-rest/src/domain/util"

// Student Entity.
type Student struct {
	Person      Person      `db:"person" json:"person,omitempty"`
	SchoolClass SchoolClass `db:"school_class" json:"school_class,omitempty"`
}

// ToString ...
func (entity *Student) ToString() string {
	result, _ := getJSONSerilizer(entity)
	return result
}

// GetTableName ...
func (entity *Student) GetTableName() string {
	return util.GetType(entity)
}
