package entidade

//Perfil estrutura que representa a tabela Perfil
type Perfil struct {
	IDPerfil        int8   `json:"id"`
	Perfil          string `json:"perfil" validate:"required"`
	RegistroApagado bool
}
