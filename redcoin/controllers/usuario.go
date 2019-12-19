package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"redcoin/repositorio"
	"redcoin/modelos"
)

//ListarUsuario método responsavel por listar 1 ou N Usuarios
func ListarUsuario(w http.ResponseWriter, r *http.Request) (erro error) {

	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		tU, e := repositorio.TodosUsuario()
		if e != nil {
			return e
		}

		json, _ := json.Marshal(tU)
		fmt.Fprint(w, string(json))
		return
	}

	id, e := strconv.Atoi(keys[0])
	if e != nil {
		return e
	}

	u, e := repositorio.IDUsuario(id)
	if e != nil {
		return e
	}

	json, _ := json.Marshal(u)
	fmt.Fprint(w, string(json))
	return
}

//PersistirUsuario método responsavel por adicionar ou alterar um Usuario
func PersistirUsuario(w http.ResponseWriter, r *http.Request) (erro error) {
	var usuario modelos.Usuario
	e := json.NewDecoder(r.Body).Decode(&usuario)
	if e != nil {
		return e
	}

	switch r.Method {
	case "POST":
		e = repositorio.AdicionarUsuario(usuario)
		w.WriteHeader(http.StatusCreated)
		break
	case "PUT":
		e = repositorio.AlterarUsuario(usuario)
		w.WriteHeader(http.StatusAccepted)
		break
	}

	if e != nil {
		return e
	}

	return nil
}