package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"redcoin/modelos"
	"redcoin/repositorio"
)

//AutenticarClienteApi método responsavel por validar as credenciais do cliente da API
func AutenticarClienteApi(w http.ResponseWriter, r *http.Request) (erro error) {
	var cliente modelos.ClienteApi
	e := json.NewDecoder(r.Body).Decode(&cliente)
	if e != nil {
		return e
	}

	existe, e := repositorio.AutenticarClienteApi(cliente.Usuario, cliente.Senha)
	if e != nil {
		return e
	}

	if !existe {
		return errors.New("Cliente não autorizado")
	}

	return nil
}

//PersistirClienteApi método responsavel por adicionar um Perfil
func PersistirClienteApi(w http.ResponseWriter, r *http.Request) {
	var cliente modelos.ClienteApi
	e := json.NewDecoder(r.Body).Decode(&cliente)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "POST":
		e = repositorio.AdicionarClienteApi(cliente)
		break
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"message":"Não há método implementado"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
