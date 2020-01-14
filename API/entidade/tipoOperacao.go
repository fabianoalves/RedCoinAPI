package entidade

//TipoOperacao estrutura que representa a tabela tipoOperacao
type TipoOperacao struct {
	IDTipoOperacao  int8   `json:"id"`
	Operacao        string `json:"operacao" validate:"required"`
	RegistroApagado bool
}
