package model

// Student Entity.
type Student struct {
	entityModel
	Person      Person      `json:"person,omitempty"`
	SchoolClass SchoolClass `json:"school_class,omitempty"`
}
