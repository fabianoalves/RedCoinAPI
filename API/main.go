package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	cn "github.com/rteles86/RedCoinApi/API/configuracoes/conexoes"
	u "github.com/rteles86/RedCoinApi/API/configuracoes/utils"
	c "github.com/rteles86/RedCoinApi/API/controllers"
	repo "github.com/rteles86/RedCoinApi/API/repositorio"
	srv "github.com/rteles86/RedCoinApi/API/servico"
)

func main() {
	fmt.Println("RedCoinApi")

	conexaoBanco := repo.NovaConexao(cn.DriverMysql, u.DecriptografarChaves(cn.ConexaoMysql))

	srv.New(conexaoBanco)

	defer conexaoBanco.Db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/autenticacao", c.AutenticacaoCliente).Methods(http.MethodPost)
	r.HandleFunc("/autenticacao/adicionar", c.PersistirCliente).Methods(http.MethodPost)

	aut := mux.NewRouter()
	aut.HandleFunc("/api/perfil/todos", c.TodosPerfil).Methods(http.MethodGet)
	r.Handle("/api/perfil/todos", commonMiddleware(aut))
	aut.HandleFunc("/api/perfil/id", c.IDPerfil).Methods(http.MethodGet)
	r.Handle("/api/perfil/id", commonMiddleware(aut))
	aut.HandleFunc("/api/perfil/adicionar", c.AdicionarPerfil).Methods(http.MethodPost)
	r.Handle("/api/perfil/adicionar", aut)
	aut.HandleFunc("/api/perfil/alterar", c.AlterarPerfil).Methods(http.MethodPut)
	r.Handle("/api/perfil/alterar", aut)

	aut.HandleFunc("/api/tipo-operacao/todos", c.TodosTipoOperacao).Methods(http.MethodGet)
	r.Handle("/api/tipo-operacao/todos", commonMiddleware(aut))
	aut.HandleFunc("/api/tipo-operacao/id", c.IDTipoOperacao).Methods(http.MethodGet)
	r.Handle("/api/tipo-operacao/id", commonMiddleware(aut))
	aut.HandleFunc("/api/tipo-operacao/adicionar", c.AdicionarTipoOperacao).Methods(http.MethodPost)
	r.Handle("/api/tipo-operacao/adicionar", commonMiddleware(aut))
	aut.HandleFunc("/api/tipo-operacao/alterar", c.AlterarTipoOperacao).Methods(http.MethodPut)
	r.Handle("/api/tipo-operacao/alterar", commonMiddleware(aut))

	aut.HandleFunc("/api/usuario/todos", c.TodosUsuario).Methods(http.MethodGet)
	r.Handle("/api/usuario/todos", commonMiddleware(aut))
	aut.HandleFunc("/api/usuario/id", c.IDUsuario).Methods(http.MethodGet)
	r.Handle("/api/usuario/id", commonMiddleware(aut))
	aut.HandleFunc("/api/usuario/adicionar", c.AdicionarUsuario).Methods(http.MethodPost)
	r.Handle("/api/usuario/adicionar", commonMiddleware(aut))
	aut.HandleFunc("/api/usuario/alterar", c.AlterarUsuario).Methods(http.MethodPut)
	r.Handle("/api/usuario/alterar", commonMiddleware(aut))

	aut.HandleFunc("/api/operacao/email", c.EmailUsuarioOperacao).Methods(http.MethodGet)
	r.Handle("/api/operacao/email", commonMiddleware(aut))
	aut.HandleFunc("/api/operacao/data", c.PeriodoOperacao).Methods(http.MethodGet)
	r.Handle("/api/operacao/data", commonMiddleware(aut))
	aut.HandleFunc("/api/operacao/gravar", c.PersistirOperacao).Methods(http.MethodPost)
	r.Handle("/api/operacao/gravar", commonMiddleware(aut))

	log.Fatal(http.ListenAndServe(":2801", r))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		if !c.ValidarTokenCliente(w, r) {
			return
		}
		next.ServeHTTP(w, r)
	})
}
