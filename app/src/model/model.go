package model

import (
	"encoding/json"
	"log"

)

type entityModel struct {
}

// Decoder - Convert JSON to structure.
func (entity *entityModel) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entity); err != nil {
		return err
	}
	return nil
}

// Encoder - Convert structure to JSON.
func (entity *entityModel) Encoder() ([]byte, error) {
	result, err := json.Marshal(entity)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return result, nil
}

// ToString ...
func (entity *entityModel) ToString() (string, error) {
	result, err := entity.Encoder()
	if err != nil {
		return "", err
	}

	return string(result), nil
}
