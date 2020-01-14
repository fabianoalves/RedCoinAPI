package infra

import (
	"encoding/json"
	"net/http"
	"net/url"

	cn "github.com/rteles86/RedCoinApi/API/configuracoes/conexoes"
	u "github.com/rteles86/RedCoinApi/API/configuracoes/utils"
	e "github.com/rteles86/RedCoinApi/API/entidade"
)

func consultaPrecoBitcoin(target interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", u.DecriptografarChaves(cn.EnderecoBitcoin), nil)
	if err != nil {
		return err
	}

	q := url.Values{}
	q.Add("id", cn.IDBitcoin)
	q.Add("amount", cn.QuantidadeBitcoin)
	q.Add("convert", cn.FormatoMoeda)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", u.DecriptografarChaves(cn.TokenBitcoin))
	req.URL.RawQuery = q.Encode()

	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

//CotacaoBitCoin Retorna o valor em Reais do total de BitCoin informado
func CotacaoBitCoin() (float64, error) {
	priceData := e.ResultBitCoin{}
	err := consultaPrecoBitcoin(&priceData)
	if err != nil {
		return 0.0, err
	}

	return priceData.Data.Quote.Brl.Price, nil
}
