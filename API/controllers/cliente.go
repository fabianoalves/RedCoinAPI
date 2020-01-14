package controllers

import (
	"net/http"

	u "github.com/rteles86/RedCoinApi/API/configuracoes/utils"
	e "github.com/rteles86/RedCoinApi/API/entidade"
	srv "github.com/rteles86/RedCoinApi/API/servico"
)

//PersistirCliente método responsavel por adicionar um Perfil
func PersistirCliente(w http.ResponseWriter, r *http.Request) {
	var cliente e.Cliente

	//converte JSON em entidade
	if !u.ParseJSONEntidade(&cliente, w, r) {
		return
	}

	//valida os dados de entrada da entidade
	err := u.New(&cliente)
	if err != nil {
		u.RespostaUnprocessableEntity(err.Error(), w)
		return
	}

	//executa a chamada do servico
	err = srv.CriarCliente(cliente)
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}
}

// AutenticacaoCliente método que realiza a autenticação da API
func AutenticacaoCliente(w http.ResponseWriter, r *http.Request) {
	var cliente e.Cliente

	//converte JSON em entidade
	if !u.ParseJSONEntidade(&cliente, w, r) {
		return
	}

	//valida os dados de entrada da entidade
	err := u.New(&cliente)
	if err != nil {
		u.RespostaUnprocessableEntity(err.Error(), w)
		return
	}

	existe, token := srv.AutenticacaoCliente(cliente)
	if !existe {
		u.RespostaStatusUnauthorized(w)
		return
	}

	mensagemBody := `{"token":"` + token + `"}`
	u.RespostaOK(w, mensagemBody)

}

//ValidarTokenCliente método que realiza a verificação do Token do cliente
func ValidarTokenCliente(w http.ResponseWriter, r *http.Request) (tokenValido bool) {
	tokenCliente := r.Header.Get("token")

	tokenValido = srv.ValidarTokenCliente(tokenCliente)
	if !tokenValido {
		u.RespostaStatusNotAcceptable(w)
	}

	return tokenValido
}
