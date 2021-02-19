package helper

import (
	"encoding/json"
	"net/http"
	"reflect"

)

// HTTPResponseOK ...
func HTTPResponseOK(w http.ResponseWriter, value interface{}) {
	_jsonResponse(w, value, 200)
}

// HTTPResponseCreate ...
func HTTPResponseCreate(w http.ResponseWriter, value interface{}) {
	_jsonResponse(w, value, 201)
}

func _jsonResponse(w http.ResponseWriter, value interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	if _isInterfaceNil(value) {
		json.NewEncoder(w)
	} else {
		json.NewEncoder(w).Encode(value)
	}
}

func _isInterfaceNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil() || i == ""
}
