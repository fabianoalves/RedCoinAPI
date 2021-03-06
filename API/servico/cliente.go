package servico

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	u "github.com/rteles86/RedCoinApi/API/configuracoes/utils"
	e "github.com/rteles86/RedCoinApi/API/entidade"
	repo "github.com/rteles86/RedCoinApi/API/repositorio"
)

//CriarCliente registra as credenciais de um cliente
func CriarCliente(cliente e.Cliente) (erro error) {
	cliente.Senha = u.CriptografarSenha(cliente.Senha)
	erro = repo.AdicionarCliente(cn, cliente)
	return erro
}

//AutenticacaoCliente verifica as credenciais de um cliente
func AutenticacaoCliente(cliente e.Cliente) (existe bool, msgToken string) {

	c := repo.Cliente(cn, cliente.Usuario)
	if !u.VerificarSenha(cliente.Senha, c.Senha) {
		return false, "Ops... usuario não tem acesso a API"
	}

	expirationTime := time.Now().Add(50 * time.Minute)

	cliente.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	claims := &cliente

	gerarToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := gerarToken.SignedString([]byte("redcoinApi2020!@"))

	return true, tokenString
}

//ValidarTokenCliente verifica se o token da API continua válido
func ValidarTokenCliente(tokenCliente string) (valido bool) {
	claims := &e.Cliente{}

	tkn, err := jwt.ParseWithClaims(tokenCliente, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("redcoinApi2020!@"), nil
	})
	if err != nil || err == jwt.ErrSignatureInvalid || !tkn.Valid {
		return false
	}

	return true
}
