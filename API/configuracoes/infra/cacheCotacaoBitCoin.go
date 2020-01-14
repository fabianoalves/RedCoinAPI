package infra

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
	cn "github.com/rteles86/RedCoinApi/API/configuracoes/conexoes"
	u "github.com/rteles86/RedCoinApi/API/configuracoes/utils"
)

func conexaoRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     u.DecriptografarChaves(cn.EnderecoRedis),
		Password: u.DecriptografarChaves(cn.SenhaRedis),
		DB:       cn.IDBancoRedis,
	})
}

//ExisteCotacao retorna se existe ou não cotação armazenada em cache
func ExisteCotacao() bool {
	client := conexaoRedis()

	return client.Exists("cotacaoBitCoin").Val() > 0
}

//CotacaoEmCache retorna o valor do Bitcoin consultado da API nos ultimos 50 minutos
func CotacaoEmCache() float64 {

	client := conexaoRedis()

	val, _ := client.Get("cotacaoBitCoin").Result()

	valor, _ := strconv.ParseFloat(val, 64)

	return valor

}

//CotacaoGravarCache grava o valor do Bitcoin consultado da API e define o tempo de expiracao de 50 minutos
func CotacaoGravarCache(valorBitCoin float64) {

	client := conexaoRedis()

	client.Set("cotacaoBitCoin", valorBitCoin, 50*time.Minute)
}
