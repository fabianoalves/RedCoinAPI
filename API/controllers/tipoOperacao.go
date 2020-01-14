package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	u "github.com/rteles86/RedCoinApi/API/configuracoes/utils"
	e "github.com/rteles86/RedCoinApi/API/entidade"
	srv "github.com/rteles86/RedCoinApi/API/servico"
)

//TodosTipoOperacao método responsavel por listar os Tipo Operacao
func TodosTipoOperacao(w http.ResponseWriter, r *http.Request) {
	w.Header()

	tTo, err := srv.TodosTipoOperacao()
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}

	json, _ := json.Marshal(tTo)
	u.RespostaOK(w, string(json))
}

//IDTipoOperacao Retorna o objeto de Tipo Operacao de acordo com o ID informado
func IDTipoOperacao(w http.ResponseWriter, r *http.Request) {
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

	to, err := srv.IDTipoOperacao(int8(id))
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}

	json, _ := json.Marshal(to)
	u.RespostaOK(w, string(json))
}

//AdicionarTipoOperacao método responsável pela criação de um novo Tipo Operacao
func AdicionarTipoOperacao(w http.ResponseWriter, r *http.Request) {
	var tipooOperacao e.TipoOperacao
	if !u.ParseJSONEntidade(&tipooOperacao, w, r) {
		return
	}

	err := u.New(&tipooOperacao)
	if err != nil {
		u.RespostaUnprocessableEntity(err.Error(), w)
		return
	}

	err = srv.AdicionarTipoOperacao(tipooOperacao)
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}

	u.RespostaOK(w, "Tipo Operacao criado com sucesso")
}

//AlterarTipoOperacao método responsável pela alteração de um Tipo Operacao
func AlterarTipoOperacao(w http.ResponseWriter, r *http.Request) {
	var tipooOperacao e.TipoOperacao
	if !u.ParseJSONEntidade(&tipooOperacao, w, r) {
		return
	}

	err := u.New(&tipooOperacao)
	if err != nil {
		u.RespostaUnprocessableEntity(err.Error(), w)
		return
	}

	err = srv.AtualizarTipoOperacao(tipooOperacao)
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}

	u.RespostaOK(w, "Tipo Operacao alterado com sucesso")
}
