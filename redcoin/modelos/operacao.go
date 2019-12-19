package modelos

import (
	"time"
)

//Operacao estrutura que representa a tabela Operacao
type Operacao struct {
	IdOperacao   int
	TipoOperacao TipoOperacao
	Vendedor     Usuario
	Comprador    Usuario
	DataOperacao time.Time
	ValorMoeda   float64
	ValorBitCoin float64
}

//UsuarioOperacao representa a saida para o cabeçalho do relatorio
type UsuarioOperacao struct {
	IdUsuario    int
	Email        string
	NomeCompleto string
	Operacoes    []Operacoes
}

//Operacoes representa as operações de compra do usuario para o relatorio
type Operacoes struct {
	IdOperacao           int
	TipoOperacao         string
	IdVendedor           int
	EmailVendedor        string
	NomeCompletoVendedor string
	DataOperacao         time.Time
	ValorMoeda           float64
	ValorBitCoin         float64
}
