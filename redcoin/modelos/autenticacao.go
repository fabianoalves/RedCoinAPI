package modelos

import (
	"github.com/dgrijalva/jwt-go"
)

//UsuariosApi representa os usuarios com acesso a API
var UsuariosApi = map[string]string{
	"master": "master",
	"client": "client",
}

//Credenciais estrutura que representa as credenciais do cliente da API
type Credenciais struct {
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
}

//Claims estrutura com as permissoes do usuario autenticado
type Claims struct {
	Usuario string `json:"usuario"`
	jwt.StandardClaims
}
