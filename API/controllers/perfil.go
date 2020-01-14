package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	u "github.com/rteles86/RedCoinApi/API/configuracoes/utils"
	e "github.com/rteles86/RedCoinApi/API/entidade"
	srv "github.com/rteles86/RedCoinApi/API/servico"
)

//TodosPerfil método responsavel por listar os Perfil
func TodosPerfil(w http.ResponseWriter, r *http.Request) {

	tP, err := srv.TodosPerfil()
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}

	json, _ := json.Marshal(tP)
	u.RespostaOK(w, string(json))
}

//IDPerfil Retorna o objeto de perfil de acordo com o ID informado
func IDPerfil(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok {
		u.RespostaBadRequest("Informe o parâmetro ID", w)
		return
	}

	id, err := strconv.Atoi(keys[0])
	if err != nil {
		u.RespostaBadRequest("Formato do ID é Inválido. Informe apenas números inteiros", w)
		return
	}

	p, err := srv.IDPerfil(int8(id))
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}

	json, _ := json.Marshal(p)
	u.RespostaOK(w, string(json))
}

//AdicionarPerfil método responsável pela criação de um novo Perfil
func AdicionarPerfil(w http.ResponseWriter, r *http.Request) {
	var perfil e.Perfil

	if !u.ParseJSONEntidade(&perfil, w, r) {
		return
	}

	err := u.New(&perfil)
	if err != nil {
		u.RespostaUnprocessableEntity(err.Error(), w)
		return
	}

	err = srv.AdicionarPerfil(perfil)
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}

	u.RespostaOK(w, "Perfil criado com sucesso")
}

//AlterarPerfil método responsável pela alteração de um Perfil
func AlterarPerfil(w http.ResponseWriter, r *http.Request) {
	var perfil e.Perfil
	if !u.ParseJSONEntidade(&perfil, w, r) {
		return
	}

	err := u.New(&perfil)
	if err != nil {
		u.RespostaUnprocessableEntity(err.Error(), w)
		return
	}

	err = srv.AtualizarPerfil(perfil)
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}

	u.RespostaOK(w, "Perfil atualizado com sucesso")
}
