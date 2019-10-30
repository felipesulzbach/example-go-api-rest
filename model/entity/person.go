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
func (entity *Person) New(id int64, name string, cpf string, cellPhone string, city string, zipCode string, address string, registrationDate time.Time) {
  *entity = Person{id, name, cpf, cellPhone, city, zipCode, address, registrationDate}
}

// Decoder - Decodes JSON for structure.
func (entity *Person) Decoder(jsonStream string) error {
  if err := json.Unmarshal([]byte(jsonStream), &entity); err != nil {
    return err
  }
  return nil
}

// ToString - Returns string with Person information.
func (entity *Person) ToString() string {
  campos := map[string]string{
    "ID":               strconv.FormatInt(entity.ID, 10),
    "Name":             entity.Name,
    "Cpf":              entity.Cpf,
    "CellPhone":        entity.CellPhone,
    "City":             entity.City,
    "ZipCode":          entity.ZipCode,
    "Address":          entity.Address,
    "RegistrationDate": util.FormatDateTime(entity.RegistrationDate),
  }
  retorno := ToString("Person", campos)
  return retorno
}
