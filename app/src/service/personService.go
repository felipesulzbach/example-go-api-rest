package service

import (
	"log"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/repository"

)

// FindAllPerson ...
func FindAllPerson() ([]*model.Person, error) {
	result, err := repository.FindAllPerson()
	if err != nil {
		return nil, err
	}

	for _, item := range result {
		log.Println(item.ToString())
	}

	return result, nil
}
