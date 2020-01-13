package entidade

import "github.com/dgrijalva/jwt-go"

//Cliente representa a tabela de usuarios que podem acessar a API
type Cliente struct {
	Usuario string `json:"usuario" validate:"required"`
	Senha   string `json:"senha" validate:"required,min=8"`
	jwt.StandardClaims
}
