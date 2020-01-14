package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	e "github.com/rteles86/RedCoinApi/API/entidade"
	"github.com/rteles86/RedCoinApi/API/servico"
)

//TodosUsuario método responsavel por listar os Usuarios
func TodosUsuario(w http.ResponseWriter, r *http.Request) {
	tU, e := servico.TodosUsuario()
	if e != nil {
		w.Write([]byte(`{"msg":"` + e.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json, _ := json.Marshal(tU)
	fmt.Fprint(w, string(json))
	return
}

//IDUsuario Retorna o objeto de Usuario de acordo com o ID informado
func IDUsuario(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	id, e := strconv.Atoi(keys[0])

	u, e := servico.IDUsuario(id)
	if e != nil {
		w.Write([]byte(`{"msg":"` + e.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json, _ := json.Marshal(u)
	fmt.Fprint(w, string(json))
	return
}

//AdicionarUsuario método responsável pela criação de um novo Usuario
func AdicionarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario e.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = e.New(&usuario)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		return
	}

	err = servico.AdicionarUsuario(usuario)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte(`{"msg":"Usuario criado com sucesso"}`))
	w.WriteHeader(http.StatusOK)
}

//AlterarUsuario método responsável pela alteração de um Usuario
func AlterarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario e.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = e.New(&usuario)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		return
	}

	err = servico.AtualizarUsuario(usuario)
	if err != nil {
		w.Write([]byte(`{"msg":"` + err.Error() + `"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte(`{"msg":"Usuario atualizado com sucesso"}`))
	w.WriteHeader(http.StatusOK)
}
