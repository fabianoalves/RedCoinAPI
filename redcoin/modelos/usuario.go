package modelos

import (
	"time"
)

//Usuario estrutura que representa a tabela Usuario
type Usuario struct {
	IdUsuario       int 
	Email           string
	Senha           string 
	Nome            string
	UltimoNome      string
	DataNascimento  time.Time 
	QuantidadeMoeda float64 
	RegistroApagado bool 
	PerfilUsuario   []Perfil
}