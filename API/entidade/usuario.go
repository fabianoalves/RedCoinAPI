package entidade

import "time"

//Usuario estrutura que representa a tabela Usuario
type Usuario struct {
	IDUsuario       int       `json:"id"`
	Email           string    `json:"email" validate:"required,email"`
	Senha           string    `json:"senha" validate:"required,min=8"`
	Nome            string    `json:"nome" validate:"required"`
	UltimoNome      string    `json:"ultimoNome" validate:"required"`
	DataNascimento  time.Time `json:"nascimento" validate:"required"`
	QuantidadeMoeda float64   `json:"quantidadeMoeda"`
	RegistroApagado bool
	PerfilUsuario   []Perfil `json:"perfil" validate:"required"`
}
