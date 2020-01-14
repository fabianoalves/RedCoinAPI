package utils

import (
	"errors"
	"time"

	"gopkg.in/go-playground/validator.v10"
)

var valida *validator.Validate

//New retorna a estrutura da entidade após sua validação dos dados
func New(entidade interface{}) error {
	return validarEntidade(entidade)
}

func validarEntidade(entidade interface{}) error {
	valida = validator.New()
	erro := valida.Struct(entidade)

	if erro != nil {
		return errors.New("Erro nos dados de entrada do método! Verifique os dados do campo: " + erro.(validator.ValidationErrors)[0].Field())
	}
	return nil
}

//ValidarEmail verifica se o Email é válido
func ValidarEmail(email string) error {
	valida = validator.New()
	erro := valida.Var(email, "required,email")

	return erro
}

//ValidarData verifica se o formato da data é válido
func ValidarData(data time.Time) error {
	valida = validator.New()
	erro := valida.Var(data, "required")

	return erro
}
