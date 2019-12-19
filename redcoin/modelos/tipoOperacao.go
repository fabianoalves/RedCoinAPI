package modelos

//TipoOperacao estrutura que representa a tabela tipoOperacao
type TipoOperacao struct {
	IDTipoOperacao  int8
	Operacao        string
	RegistroApagado bool
}
