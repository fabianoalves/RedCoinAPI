package main

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"redcoin/controllers"
)

func get(w http.ResponseWriter, r *http.Request) {
	var e error
	w.Header().Set("Content-Type", "application/json")

	switch {
	case strings.Contains(r.URL.Path, "TipoOperacao"):
		e = controllers.TodosTipoOperacao(w, r)
		break
	case strings.Contains(r.URL.Path, "Perfil"):
		e = controllers.TodosPerfil(w, r)
		break
	case strings.Contains(r.URL.Path, "Usuario"):
		e = controllers.ListarUsuario(w, r)
		break
	case strings.Contains(r.URL.Path, "Operacao/Email"):
		e = controllers.EmailUsuarioOperacao(w,r)
		break
	case strings.Contains(r.URL.Path, "Operacao/Data"):
		e = controllers.PeriodoOperacao(w,r)
		break
	default:
		e = errors.New("PATH NOT FOUND")
		break
	}

	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"msg":"` + e.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func postPut(w http.ResponseWriter, r *http.Request) {
	var e error
	w.Header().Set("Content-Type", "application/json")

	switch {
	case strings.Contains(r.URL.Path, "TipoOperacao"):
		e = controllers.PersistirTipoOperacao(w, r)
		break
	case strings.Contains(r.URL.Path, "Perfil"):
		e = controllers.PersistirPerfil(w, r)
		break
	case strings.Contains(r.URL.Path, "Usuario"):
		e = controllers.PersistirUsuario(w, r)
		break
	case strings.Contains(r.URL.Path, "Operacao"):
		e = controllers.PersistirOperacao(w,r)
		break
	default:
		e = errors.New("PATH NOT FOUND")
		break
	}

	if e != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"msg":"` + e.Error() + `"}`))
		return
	}

	w.Write([]byte(`{"msg":"Ok"}`))
}

func urlChamada(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		get(w, r)
	case "POST", "PUT":
		postPut(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"message":"Não há método implementado"}`))
	}
}

//Index home da Api
func Index() {
	http.HandleFunc("/api/v1/TipoOperacao/", urlChamada)
	http.HandleFunc("/api/v1/Perfil/", urlChamada)
	http.HandleFunc("/api/v1/Usuario/", urlChamada)
	http.HandleFunc("/api/v1/Operacao/", urlChamada)
	http.HandleFunc("/api/v1/Operacao/Email/", urlChamada)
	http.HandleFunc("/api/v1/Operacao/Data/", urlChamada)
	
	log.Fatal(http.ListenAndServe(":8086", nil))
}

func main() {
	Index()
}
