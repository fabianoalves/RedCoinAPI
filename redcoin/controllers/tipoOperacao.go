package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"redcoin/repositorio"
	"redcoin/modelos"
)

//TodosTipoOperacao método responsavel por listar os TipoOperacao
func TodosTipoOperacao(w http.ResponseWriter, r *http.Request) (erro error) {

	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		tTO, e := repositorio.TodosTipoOperacao()
		if e != nil {
			return e
		}

		json, _ := json.Marshal(tTO)
		fmt.Fprint(w, string(json))
		return
	}
	id, e := strconv.Atoi(keys[0])
	if e != nil {
		return e
	}

	tO, e := repositorio.IDTipoOperacao(int8(id))
	if e != nil {
		return e
	}

	json, _ := json.Marshal(tO)
	fmt.Fprint(w, string(json))
	return
}

//PersistirTipoOperacao método responsavel por adicionar um TipoOperacao
func PersistirTipoOperacao(w http.ResponseWriter, r *http.Request) (erro error) {
	var tipoOperacao modelos.TipoOperacao
	e := json.NewDecoder(r.Body).Decode(&tipoOperacao)
	if e != nil {
		return e
	}

	switch r.Method {
	case "POST":
		e = repositorio.AdicionarTipoOperacao(tipoOperacao)
		break
	case "PUT":
		e = repositorio.AlterarTipoOperacao(tipoOperacao)
		break
	}

	if e != nil {
		return e
	}

	return nil
}
