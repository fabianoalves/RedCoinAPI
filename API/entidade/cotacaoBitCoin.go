package entidade

import "time"

//ResultBitCoin Estrutura para API de BitCoin
type ResultBitCoin struct {
	Status Status
	Data   Data
}

//Status Estrutura para API de BitCoin
type Status struct {
	Timestamp    time.Time
	ErrorCode    int    `json:"Error_code"`
	ErrorMessage string `json:"Error_message"`
	Elapsed      int
	CreditCount  int `json:"Credit_count"`
	Notice       string
}

//Data Estrutura para API de BitCoin
type Data struct {
	ID         int `json:"Id"`
	Symbol     string
	Name       string
	Amount     float64
	LastUpdate time.Time `json:"Last_update"`
	Quote      Quote
}

//Quote Estrutura para API de BitCoin
type Quote struct {
	Brl BRL
}

//BRL Estrutura para API de BitCoin
type BRL struct {
	Price      float64
	LastUpdate time.Time `json:"Last_update"`
}
