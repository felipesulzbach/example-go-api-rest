package util

import (
	"encoding/json"

)

// Unserializer - Convert JSON to structure.
func Unserializer(value []byte, entity interface{}) error {
	if err := json.Unmarshal(value, &entity); err != nil {
		return err
	}

	return nil
}

// Serializer - Convert structure to JSON.
func Serializer(entity interface{}) ([]byte, error) {
	result, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	return result, nil
}
