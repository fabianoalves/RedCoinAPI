package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	e "github.com/rteles86/RedCoinApi/redcoin/entidade"
	"github.com/rteles86/RedCoinApi/redcoin/servico"
)

//TodosPerfil método responsavel por listar os Perfil
func TodosPerfil(w http.ResponseWriter, r *http.Request) {
	w.Header()

	tP, e := servico.TodosPerfil()
	if e != nil {
		w.Write([]byte(`{"msg":"` + e.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json, _ := json.Marshal(tP)
	fmt.Fprint(w, string(json))
	return
}

//IDPerfil Retorna o objeto de perfil de acordo com o ID informado
func IDPerfil(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	id, e := strconv.Atoi(keys[0])

	p, e := servico.IDPerfil(int8(id))
	if e != nil {
		w.Write([]byte(`{"msg":"` + e.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json, _ := json.Marshal(p)
	fmt.Fprint(w, string(json))
	return
}

//AdicionarPerfil método responsável pela criação de um novo Perfil
func AdicionarPerfil(w http.ResponseWriter, r *http.Request) {
	var perfil e.Perfil
	err := json.NewDecoder(r.Body).Decode(&perfil)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = e.New(&perfil)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		return
	}

	err = servico.AdicionarPerfil(perfil)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte(`{"msg":"Perfil criado com sucesso"}`))
	w.WriteHeader(http.StatusOK)
}

//AlterarPerfil método responsável pela alteração de um Perfil
func AlterarPerfil(w http.ResponseWriter, r *http.Request) {
	var perfil e.Perfil
	err := json.NewDecoder(r.Body).Decode(&perfil)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = e.New(&perfil)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		return
	}

	err = servico.AtualizarPerfil(perfil)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte(`{"msg":"Perfil atualizado com sucesso"}`))
	w.WriteHeader(http.StatusOK)
}
