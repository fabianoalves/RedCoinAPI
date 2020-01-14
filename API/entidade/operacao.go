package entidade

import "time"

//NovaOperacao estrutura que representa a tabela Operacao
type NovaOperacao struct {
	IDOperacao     int       `json:"id"`
	IDTipoOperacao int       `json:"idTipo" validate:"required,min=1"`
	IDVendedor     int       `json:"idVendedor" validate:"required,min=1"`
	IDComprador    int       `json:"idComprador" validate:"required,min=1"`
	DataOperacao   time.Time `json:"data" validate:"required"`
	ValorMoeda     float64   `json:"valorMoeda" validate:"required"`
	ValorBitCoin   float64   `json:"valorBitcoin" validate:"required"`
}

//Operacao representa a saida do relatorio por Email
type Operacao struct {
	IDOperacao   int          `json:"id"`
	TipoOperacao TipoOperacao `json:"tipo"`
	Vendedor     Usuario      `json:"vendedor"`
	Comprador    Usuario      `json:"comprador"`
	DataOperacao time.Time    `json:"data"`
	ValorMoeda   float64      `json:"valorMoeda"`
	ValorBitCoin float64      `json:"valorBitcoin"`
}

//UsuarioOperacao representa a saida para o cabeçalho do relatorio por Periodo
type UsuarioOperacao struct {
	IDUsuario    int         `json:"idUsuario"`
	Email        string      `json:"email"`
	NomeCompleto string      `json:"nome"`
	Operacoes    []Operacoes `json:"operacoes"`
}

//Operacoes representa as operações de compra do usuario para o relatorio por periodo
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
