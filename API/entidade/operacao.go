package entidade

import (
	"time"
)

//Operacao estrutura que representa a tabela Operacao
type Operacao struct {
	IDOperacao   int          `json:"id"`
	TipoOperacao TipoOperacao `json:"tipo"`
	Vendedor     Usuario      `json:"vendedor"`
	Comprador    Usuario      `json:"comprador"`
	DataOperacao time.Time    `json:"data"`
	ValorMoeda   float64      `json:"valorMoeda"`
	ValorBitCoin float64      `json:"valorBitcoin"`
}

//UsuarioOperacao representa a saida para o cabeçalho do relatorio
type UsuarioOperacao struct {
	IDUsuario    int         `json:"idUsuario"`
	Email        string      `json:"email"`
	NomeCompleto string      `json:"nome"`
	Operacoes    []Operacoes `json:"operacoes"`
}

//Operacoes representa as operações de compra do usuario para o relatorio
type Operacoes struct {
	IDOperacao           int       `json:"id"`
	TipoOperacao         string    `json:"tipo"`
	IDVendedor           int       `json:"idVendedor"`
	EmailVendedor        string    `json:"emailVendedor"`
	NomeCompletoVendedor string    `json:"nomeVendedor"`
	DataOperacao         time.Time `json:"data"`
	ValorMoeda           float64   `json:"valorMoeda"`
	ValorBitCoin         float64   `json:"valorBitcoin"`
}
