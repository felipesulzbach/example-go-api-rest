package model

// Teacher Entity.
type Teacher struct {
	entityModel
	Person Person `json:"person,omitempty"`
	Course Course `json:"course,omitempty"`
}
