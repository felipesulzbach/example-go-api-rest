package entity

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/_dev/exemplo-api-rest/util"
)

// Person Entity.
type Person struct {
	ID               int64     `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Cpf              string    `json:"cpf,omitempty"`
	CellPhone        string    `json:"cellPhone,omitempty"`
	City             string    `json:"city,omitempty"`
	ZipCode          string    `json:"zipCode,omitempty"`
	Address          string    `json:"address,omitempty"`
	RegistrationDate time.Time `json:"registrationDate,omitempty"`
}

// New - Loads a new Person structure.
func (entidade *Person) New(id int64, name string, cpf string, cellPhone string, city string, zipCode string, address string, registrationDate time.Time) {
	*entidade = Person{id, name, cpf, cellPhone, city, zipCode, address, registrationDate}
}

// Decoder - Decodes JSON for structure.
func (entidade *Person) Decoder(jsonStream string) error {
	if err := json.Unmarshal([]byte(jsonStream), &entidade); err != nil {
		return err
	}
	return nil
}

// ToString - Returns string with Person information.
func (entidade *Person) ToString() string {
	campos := map[string]string{
		"ID":               strconv.FormatInt(entidade.ID, 10),
		"Name":             entidade.Name,
		"Cpf":              entidade.Cpf,
		"CellPhone":        entidade.CellPhone,
		"City":             entidade.City,
		"ZipCode":          entidade.ZipCode,
		"Address":          entidade.Address,
		"RegistrationDate": util.FormatarDataHora(entidade.RegistrationDate),
	}
	retorno := ToString("Person", campos)
	return retorno
}
