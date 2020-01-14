package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	e "github.com/rteles86/RedCoinApi/API/entidade"
	"github.com/rteles86/RedCoinApi/API/servico"
)

//TodosTipoOperacao método responsavel por listar os Tipo Operacao
func TodosTipoOperacao(w http.ResponseWriter, r *http.Request) {
	w.Header()

	tTo, e := servico.TodosTipoOperacao()
	if e != nil {
		w.Write([]byte(`{"msg":"` + e.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json, _ := json.Marshal(tTo)
	fmt.Fprint(w, string(json))
	return
}

//IDTipoOperacao Retorna o objeto de Tipo Operacao de acordo com o ID informado
func IDTipoOperacao(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	id, e := strconv.Atoi(keys[0])

	to, e := servico.IDTipoOperacao(int8(id))
	if e != nil {
		w.Write([]byte(`{"msg":"` + e.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json, _ := json.Marshal(to)
	fmt.Fprint(w, string(json))
	return
}

//AdicionarTipoOperacao método responsável pela criação de um novo Tipo Operacao
func AdicionarTipoOperacao(w http.ResponseWriter, r *http.Request) {
	var tipooOperacao e.TipoOperacao
	err := json.NewDecoder(r.Body).Decode(&tipooOperacao)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = e.New(&tipooOperacao)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		return
	}

	err = servico.AdicionarTipoOperacao(tipooOperacao)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//AlterarTipoOperacao método responsável pela alteração de um Tipo Operacao
func AlterarTipoOperacao(w http.ResponseWriter, r *http.Request) {
	var tipooOperacao e.TipoOperacao
	err := json.NewDecoder(r.Body).Decode(&tipooOperacao)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = e.New(&tipooOperacao)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		return
	}

	err = servico.AtualizarTipoOperacao(tipooOperacao)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte(`{"msg":"TipoOperacao criado com sucesso"}`))
	w.WriteHeader(http.StatusOK)
}
