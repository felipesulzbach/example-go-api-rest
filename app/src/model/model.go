package model

import (
	"log"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/util"

)

func getJSONSerilizer(entity interface{}) (string, error) {
	result, err := util.Serializer(entity)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return string(result), nil
}
