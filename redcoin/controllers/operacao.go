package controllers

import (
	"encoding/json"
	"net/http"
	"errors"
	"fmt"
	"redcoin/modelos"
	"redcoin/repositorio"
	"time"
)

//EmailUsuarioOperacao método responsavel por listar as operações de um determinado usuario através de seu Email
func EmailUsuarioOperacao(w http.ResponseWriter, r *http.Request) (erro error) {

	keys, ok := r.URL.Query()["email"]
	if !ok || len(keys[0]) < 1 {
			return errors.New("Necessário parâmetro email")
	}

	email := keys[0]

	eOp, e := repositorio.EmailUsuarioOperacao(email)
	if e != nil {
		return e
	}

	json, _ := json.Marshal(eOp)
	fmt.Fprint(w, string(json))
	return
}

//PeriodoOperacao método responsavel por listar as operações de uma determinada Data
func PeriodoOperacao(w http.ResponseWriter, r *http.Request) (erro error) {

	keys, ok := r.URL.Query()["data"]
	if !ok || len(keys[0]) < 1 {
			return errors.New(`Necessário parâmetro Data no Formato "YYYY-MES-dd" exemplo: "2019-FEB-01"`)
	}

	sdata := keys[0]
	data, e := time.Parse("2006-01-02", sdata)

	pOp, e := repositorio.PeriodoOperacao(data)
	if e != nil {
		return e
	}

	json, _ := json.Marshal(pOp)
	fmt.Fprint(w, string(json))
	return
}

//PersistirOperacao método responsavel por adicionar uma operaçao
func PersistirOperacao(w http.ResponseWriter, r *http.Request) (erro error) {
	var operacao modelos.Operacao
	var existeCache bool

	e := json.NewDecoder(r.Body).Decode(&operacao)
	if e != nil {
		return e
	}

	operacao.ValorMoeda, existeCache = repositorio.CotacaoEmCache()

	if existeCache != true{
		operacao.ValorMoeda, e = repositorio.CotacaoBitCoin()
		if e != nil {
			return e
		}
		repositorio.CotacaoGravarCache(operacao.ValorMoeda)
	}
	
	operacao.ValorMoeda = operacao.ValorMoeda * operacao.ValorBitCoin

	switch r.Method {
	case "POST":
		e = repositorio.AdicionarOperacao(operacao)
		break
	}
	if e != nil {
		return e
	}

	return nil
}