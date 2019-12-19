package conexoes


//Mysql representa as credenciais de acesso ao banco
type Mysql struct {
	Driver            string
	Usuario           string
	Senha             string
	NomeBanco         string
	ConfiguracaoTempo string
	Porta             string
}

//ApiBitCoin representa as credenciais de acesso a API
type ApiBitCoin struct{
	Url string
	ID string
	Amount string
	Convert string
	Token string
}

//Redis representa as credenciais de acesso ao banco REDIS
type Redis struct{
	Endereco string
	Senha string
	IDBanco int
}

//CredenciaisMysql retorna as credenciais de acesso ao banco
func CredenciaisMysql() (credenciais Mysql) {
	var mysql Mysql
	mysql.Driver = "mysql"
	mysql.Usuario = "root"
	mysql.Senha = "123mudar"
	mysql.NomeBanco = "redcoin"
	mysql.ConfiguracaoTempo = "?parseTime=true"
	mysql.Porta = "tcp(localhost:3308)"

	return mysql
}
//CredenciaisApiBitCoin retorna as credenciais de acesso a Api de cotação
func CredenciaisApiBitCoin()(credenciais ApiBitCoin){
	var api ApiBitCoin
	api.Url = "https://pro-api.coinmarketcap.com/v1/tools/price-conversion"
	api.ID = "1"
	api.Amount = "1"
	api.Convert = "BRL"
	api.Token = "fef41ec5-5436-4440-adca-bad8e2380978"

	return api
}

func CredencialRedis()(credenciais Redis){
	var redis Redis
	redis.Endereco = "localhost:6379"
	redis.Senha = ""
	redis.IDBanco = 0

	return redis
}