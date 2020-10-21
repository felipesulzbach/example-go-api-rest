package model

// Profile Entity.
type Profile struct {
	entityModel
	ID               int64  `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	Code             string `json:"code,omitempty"`
	AccessPermission string `json:"access_permission,omitempty"`
}
