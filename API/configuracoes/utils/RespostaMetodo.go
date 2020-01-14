package utils

import "net/http"

//RespostaOK retorna o objeto response com status 200
func RespostaOK(w http.ResponseWriter, mensagemBody string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(mensagemBody))
}

//RespostaInternalServerError retorna o objeto response com status 500
func RespostaInternalServerError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Ocorreu um problema interno na chamada do método"))
}

//RespostaBadRequest retorna o objeto response com status 400
func RespostaBadRequest(mensagem string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(mensagem))
}

//RespostaUnprocessableEntity retorna o objeto response com status 422
func RespostaUnprocessableEntity(mensagem string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte(mensagem))
}

//RespostaStatusNotFound retorna o objeto response com status 404
func RespostaStatusNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Método não foi encontrado"))
}

//RespostaStatusUnauthorized retorna o objeto response com status 401
func RespostaStatusUnauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("As credenciais utilizadas para acessar o Método são inválidas"))
}

//RespostaStatusNotAcceptable retorna o objeto response com status 406
func RespostaStatusNotAcceptable(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNotAcceptable)
	w.Write([]byte("Token inválido"))
}
