package entidade

import (
	"errors"

	"gopkg.in/go-playground/validator.v10"
)

var valida *validator.Validate

func validateStruct(c interface{}) error {
	valida = validator.New()

	erro := valida.Struct(c)

	if erro != nil {
		return errors.New(`Foi encontrado erro nos dados de entrada do método! 
		Verifique as informações do campo: ` + erro.(validator.ValidationErrors)[0].Field())
	}
	return nil
}

//New retorna a estrutura da entidade após sua validação dos dados
func New(c interface{}) error {
	return validateStruct(c)
}
