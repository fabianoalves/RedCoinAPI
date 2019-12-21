package controllers

import (
	"encoding/json"
	"net/http"
	"redcoin/conexoes"
	"redcoin/modelos"
	"redcoin/repositorio"

	"time"

	"github.com/dgrijalva/jwt-go"
)

// AutenticacaoApi método que realiza a autenticação da API
func AutenticacaoApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cliente modelos.ClienteApi

	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	existe, err := repositorio.AutenticarClienteApi(cliente.Usuario, cliente.Senha)

	if !existe {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &modelos.Claims{
		Usuario: cliente.Usuario,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(conexoes.ChaveJwt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`
	{
		"msg":"Token criado com êxito, utilize-o como parametro Header nas chamadas das demais rotas",
		"token":"` + tokenString + `"
	}
	`))
}

func ValidarToken(w http.ResponseWriter, r *http.Request) (tokenValido bool) {
	w.Header().Set("Content-Type", "application/json")

	tokenApi := r.Header.Get("token")

	claims := &modelos.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenApi, claims, func(token *jwt.Token) (interface{}, error) {
		return conexoes.ChaveJwt, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message":"token inválido!"}`))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"token inválido!"}`))
		return
	}

	return true
}
