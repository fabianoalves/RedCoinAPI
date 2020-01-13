package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	e "github.com/rteles86/RedCoinApi/redcoin/entidade"
	"github.com/rteles86/RedCoinApi/redcoin/servico"
)

//EmailUsuarioOperacao método responsavel por listar as operações de um determinado usuario através de seu Email
func EmailUsuarioOperacao(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["email"]
	if !ok || len(keys[0]) < 1 {
		w.Write([]byte(`{"msg":"Necessário parâmetro email"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	email := keys[0]

	eOp, e := servico.EmailUsuarioOperacao(email)
	if e != nil {
		w.Write([]byte(`{"msg":"` + e.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json, _ := json.Marshal(eOp)
	fmt.Fprint(w, string(json))
	return
}

//PeriodoOperacao método responsavel por listar as operações de uma determinada Data
func PeriodoOperacao(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["data"]
	if !ok || len(keys[0]) < 1 {
		w.Write([]byte(`{"msg":"Necessário parâmetro Data no Formato "YYYY-MES-dd" exemplo: "2019-FEB-01""}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sdata := keys[0]
	data, e := time.Parse("2006-01-02", sdata)

	pOp, e := servico.PeriodoOperacao(data)
	if e != nil {
		w.Write([]byte(`{"msg":"` + e.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json, _ := json.Marshal(pOp)
	fmt.Fprint(w, string(json))
	return
}

//PersistirOperacao método responsavel por adicionar uma operaçao
func PersistirOperacao(w http.ResponseWriter, r *http.Request) {
	var operacao e.Operacao
	err := json.NewDecoder(r.Body).Decode(&operacao)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = e.New(&operacao)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		return
	}

	err = servico.PersistirOperacao(operacao)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte(`{"msg":"Transação de Bitcoin criada com sucesso"}`))
	w.WriteHeader(http.StatusOK)
}
