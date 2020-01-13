package entidade

//PerfilUsuario estrutura que representa a tabela PerfilUsuario
type PerfilUsuario struct {
	IDPerfilUsuario int  `json:"id"`
	IDPerfil        int8 `json:"idPerfil" validate:"required,min=1"`
	IDUsuario       int  `json:"idUsuario" validate:"required,min=1"`
}
