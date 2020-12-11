package contract

import (
	"errors"
	"time"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"

)

/*// PersonRequest ...
type PersonRequest struct {
	ID        int64
	Name      string
	Cpf       string
	CellPhone string `validate:"nonzero,regexp=^(0|\\+62|062|62)[0-9]+$"`
	City      string
	ZipCode   string
	Address   string
}*/

// PersonContract ...
type PersonContract struct {
	Entity model.Person
}

// ValidateBodyCreate ...
func (contract PersonContract) ValidateBodyCreate() (model.Person, error) {
	var entity model.Person
	if err := contract._validateCreate(); err != nil {
		return entity, err
	}

	request := model.Person{}
	request.Name = contract.Entity.Name
	request.Cpf = contract.Entity.Cpf
	request.CellPhone = contract.Entity.CellPhone
	request.City = contract.Entity.City
	request.ZipCode = contract.Entity.ZipCode
	request.Address = contract.Entity.Address
	request.RegistrationDate = time.Now()

	return request, nil
}

// ValidateBodyUpdate ...
func (contract PersonContract) ValidateBodyUpdate() (model.Person, error) {
	var entity model.Person
	if err := contract._validateUpdate(); err != nil {
		return entity, err
	}

	request := model.Person{}
	request.Name = contract.Entity.Name
	request.Cpf = contract.Entity.Cpf
	request.CellPhone = contract.Entity.CellPhone
	request.City = contract.Entity.City
	request.ZipCode = contract.Entity.ZipCode
	request.Address = contract.Entity.Address

	return request, nil
}

func (contract PersonContract) _validateCreate() error {
	if err := _validateCpf(contract.Entity.Cpf); err != nil {
		return err
	}

	return nil
}

func (contract PersonContract) _validateUpdate() error {
	if err := validateID(contract.Entity.ID); err != nil {
		return err
	}

	if err := contract._validateCreate(); err != nil {
		return err
	}

	return nil
}

func _validateCpf(value string) error {
	if value == "" {
		return errors.New("The CPF parameter must be informed")
	}

	if err := validateStringOnlyNumbers("CPF", value); err != nil {
		return err
	}

	if len(value) != 11 {
		return errors.New("The CPF parameter must contain 11 digits")
	}

	return nil
}
