package contractcourse

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

)

// GetCreateRequestBody ...
func GetCreateRequestBody(r *http.Request) (Post, error) {
	var requestBody Post
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return requestBody, err
	}

	return requestBody, nil
}

// GetUpdateRequestBody ...
func GetUpdateRequestBody(r *http.Request) (Put, error) {
	var requestBody Put
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return requestBody, err
	}

	return requestBody, nil
}

// GetPath ...
func GetPath(r *http.Request) (Get, error) {
	var requestPath Get
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	requestPath.ID = id

	return requestPath, nil
}

// Post ...
type Post struct {
	Name        string `json:"name" validate:"alphanumeric,required"`
	Description string `json:"description" validate:"alphanumeric"`
}

// Put ...
type Put struct {
	ID          int64  `json:"id" validate:"numeric,required"`
	Name        string `json:"name" validate:"alphanumeric,required"`
	Description string `json:"description" validate:"alphanumeric"`
}

// Get ...
type Get struct {
	ID int64 `json:"id" validate:"numeric,required"`
}
