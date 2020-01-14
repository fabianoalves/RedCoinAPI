package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	u "github.com/rteles86/RedCoinApi/API/configuracoes/utils"
	e "github.com/rteles86/RedCoinApi/API/entidade"
	srv "github.com/rteles86/RedCoinApi/API/servico"
)

//TodosUsuario método responsavel por listar os Usuarios
func TodosUsuario(w http.ResponseWriter, r *http.Request) {
	tU, err := srv.TodosUsuario()
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}

	json, _ := json.Marshal(tU)
	u.RespostaOK(w, string(json))
}

//IDUsuario Retorna o objeto de Usuario de acordo com o ID informado
func IDUsuario(w http.ResponseWriter, r *http.Request) {
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

	usu, err := srv.IDUsuario(id)
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}

	json, _ := json.Marshal(usu)
	u.RespostaOK(w, string(json))
}

//AdicionarUsuario método responsável pela criação de um novo Usuario
func AdicionarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario e.Usuario
	if !u.ParseJSONEntidade(&usuario, w, r) {
		return
	}

	err := u.New(&usuario)
	if err != nil {
		u.RespostaUnprocessableEntity(err.Error(), w)
		return
	}

	err = srv.AdicionarUsuario(usuario)
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}

	u.RespostaOK(w, "Usuario criado com sucesso")
}

//AlterarUsuario método responsável pela alteração de um Usuario
func AlterarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario e.Usuario
	if !u.ParseJSONEntidade(&usuario, w, r) {
		return
	}

	err := u.New(&usuario)
	if err != nil {
		u.RespostaUnprocessableEntity(err.Error(), w)
		return
	}

	err = srv.AtualizarUsuario(usuario)
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}

	u.RespostaOK(w, "Usuario atualizado com sucesso")
}
