package modelos

import "time"

//ResultBitCoin Estrutura para API de BitCoin
type ResultBitCoin struct {
	Status Status
	Data   Data
}

//Status Estrutura para API de BitCoin
type Status struct {
	Timestamp     time.Time
	Error_code    int
	Error_message string
	Elapsed       int
	Credit_count  int
	Notice        string
}

//Data Estrutura para API de BitCoin
type Data struct {
	Id          int
	Symbol      string
	Name        string
	Amount      float64
	Last_update time.Time
	Quote       Quote
}

//Quote Estrutura para API de BitCoin
type Quote struct {
	Brl BRL
}

//BRL Estrutura para API de BitCoin
type BRL struct {
	Price       float64
	Last_update time.Time
}
