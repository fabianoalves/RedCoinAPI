package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"redcoin/repositorio"
	"redcoin/modelos"
)

//TodosPerfil método responsavel por listar os Perfil
func TodosPerfil(w http.ResponseWriter, r *http.Request) (erro error) {

	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		tP, e := repositorio.TodosPerfil()
		if e != nil {
			return e
		}

		json, _ := json.Marshal(tP)
		fmt.Fprint(w, string(json))
		return
	}

	id, e := strconv.Atoi(keys[0])
	if e != nil {
		return e
	}

	p, e := repositorio.IDPerfil(int8(id))
	if e != nil {
		return e
	}

	json, _ := json.Marshal(p)
	fmt.Fprint(w, string(json))
	return
}

//PersistirPerfil método responsavel por adicionar um Perfil
func PersistirPerfil(w http.ResponseWriter, r *http.Request) (erro error) {
	var perfil modelos.Perfil
	e := json.NewDecoder(r.Body).Decode(&perfil)
	if e != nil {
		return e
	}

	switch r.Method {
	case "POST":
		e = repositorio.AdicionarPerfil(perfil)
		break
	case "PUT":
		e = repositorio.AlterarPerfil(perfil)
		break
	}

	if e != nil {
		return e
	}

	return nil
}
