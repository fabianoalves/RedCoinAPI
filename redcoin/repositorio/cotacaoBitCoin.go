package repositorio

import(
  "encoding/json"
  "net/http"
  "net/url"

  "redcoin/modelos"
  "redcoin/conexoes"
  
)

func getJson(target interface{}) error {
	client := &http.Client{}
	apiCn := conexoes.CredenciaisApiBitCoin()

	req, err := http.NewRequest("GET",apiCn.Url, nil)
	if err != nil {
	  return err
	}
  
	q := url.Values{}
	q.Add("id", apiCn.ID)
	q.Add("amount", apiCn.Amount)
	q.Add("convert", apiCn.Convert)
  
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", apiCn.Token)
	req.URL.RawQuery = q.Encode()

	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

//CotacaoBitCoin Retorna o valor em Reais do total de BitCoin informado
func  CotacaoBitCoin()(float64,error) {
  priceData := modelos.ResultBitCoin{}
  err := getJson(&priceData)
  if err != nil {
     return 0.0, err
  }

  return priceData.Data.Quote.Brl.Price, nil
}
