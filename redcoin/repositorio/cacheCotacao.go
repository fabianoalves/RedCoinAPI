package repositorio

import(
	"redcoin/conexoes"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)


func CotacaoEmCache()(float64, bool){
	conRedis := conexoes.CredencialRedis()
	client := redis.NewClient(&redis.Options{
		Addr:     conRedis.Endereco,
		Password: conRedis.Senha,
		DB:       conRedis.IDBanco,
	})

	val, _ := client.Get("cotacaoBitCoin").Result()

	if val != "" {
		valor, e:= strconv.ParseFloat(val, 64)
		if e !=nil{
			return 0.0, false
		}
		return valor, true
	} else {
		return 0.0, false
	}
}

func CotacaoGravarCache(valorBitCoin float64){
	conRedis := conexoes.CredencialRedis()
	client := redis.NewClient(&redis.Options{
		Addr:     conRedis.Endereco,
		Password: conRedis.Senha,
		DB:       conRedis.IDBanco,
	})

	client.Set("cotacaoBitCoin", valorBitCoin, 50*time.Minute)
}