package utils

import (
	"encoding/json"
	"net/http"
)

//ParseJSONEntidade converte o JSON do corpo do método na entidade informada no método
func ParseJSONEntidade(entidade interface{}, w http.ResponseWriter, r *http.Request) bool {
	err := json.NewDecoder(r.Body).Decode(&entidade)
	if err != nil {
		RespostaBadRequest(err.Error(), w)
		return false
	}

	return true
}
