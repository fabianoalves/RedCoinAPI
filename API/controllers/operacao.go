package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	u "github.com/rteles86/RedCoinApi/API/configuracoes/utils"
	e "github.com/rteles86/RedCoinApi/API/entidade"
	"github.com/rteles86/RedCoinApi/API/servico"
	srv "github.com/rteles86/RedCoinApi/API/servico"
)

//EmailUsuarioOperacao método responsavel por listar as operações de um determinado usuario através de seu Email
func EmailUsuarioOperacao(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["email"]

	if !ok {
		u.RespostaBadRequest("Informe o parâmetro email", w)
		return
	}

	erro := u.ValidarEmail(keys[0])
	if erro != nil {
		u.RespostaUnprocessableEntity(erro.Error(), w)
		return
	}

	email := keys[0]

	eOp, e := srv.EmailUsuarioOperacao(email)
	if e != nil {
		u.RespostaInternalServerError(w)
		return
	}

	json, _ := json.Marshal(eOp)
	u.RespostaOK(w, string(json))
}

//PeriodoOperacao método responsavel por listar as operações de uma determinada Data
func PeriodoOperacao(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["data"]
	if !ok {
		u.RespostaBadRequest("Informe o parâmetro data", w)
		return
	}

	sdata := keys[0]
	data, e := time.Parse("2006-01-02", sdata)
	if e != nil {
		u.RespostaBadRequest("Formato da data é Inválido. Informe no formato YYYY-MM-DD", w)
		return
	}

	erro := u.ValidarData(data)
	if erro != nil {
		u.RespostaUnprocessableEntity(erro.Error(), w)
		return
	}

	pOp, e := servico.PeriodoOperacao(data)
	if e != nil {
		u.RespostaInternalServerError(w)
		return
	}

	json, _ := json.Marshal(pOp)
	u.RespostaOK(w, string(json))
}

//PersistirOperacao método responsavel por adicionar uma operaçao
func PersistirOperacao(w http.ResponseWriter, r *http.Request) {
	var operacao e.NovaOperacao

	//converte JSON em entidade
	if !u.ParseJSONEntidade(&operacao, w, r) {
		return
	}

	//valida os dados de entrada da entidade
	err := u.New(&operacao)
	if err != nil {
		u.RespostaUnprocessableEntity(err.Error(), w)
		return
	}

	err = servico.PersistirOperacao(operacao)
	if err != nil {
		u.RespostaInternalServerError(w)
		return
	}

	u.RespostaOK(w, "Transação de Bitcoin criada com sucesso")
}
