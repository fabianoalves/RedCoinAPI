package controllers

import (
	"encoding/json"
	"net/http"

	e "github.com/rteles86/RedCoinApi/API/entidade"
	srv "github.com/rteles86/RedCoinApi/API/servico"
)

//PersistirCliente método responsavel por adicionar um Perfil
func PersistirCliente(w http.ResponseWriter, r *http.Request) {
	var cliente e.Cliente
	e := json.NewDecoder(r.Body).Decode(&cliente)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e = srv.CriarCliente(cliente)

	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"msg":"` + e.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg":"cliente adicionado com sucesso"}`))
	return
}

// AutenticacaoCliente método que realiza a autenticação da API
func AutenticacaoCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cliente e.Cliente

	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		return
	}

	err = e.New(&cliente)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		return
	}

	existe, token, err := srv.AutenticacaoCliente(cliente)

	if !existe {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
	{
		"msg":"Token criado com êxito, utilize-o como parametro Header nas chamadas das demais rotas",
		"token":"` + token + `"
	}
	`))
}

//ValidarTokenCliente método que realiza a verificação do Token do cliente
func ValidarTokenCliente(w http.ResponseWriter, r *http.Request) (tokenValido bool) {
	w.Header().Set("Content-Type", "application/json")
	tokenCliente := r.Header.Get("token")

	tokenValido = srv.ValidarTokenCliente(tokenCliente)

	if !tokenValido {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`
		{
			"msg":"Token inválido"
		}
		`))
	}
	return tokenValido
}
