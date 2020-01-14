package repositorio

import (
	"database/sql"

	//Driver do Mysql
	_ "github.com/go-sql-driver/mysql"
)

type repoMysql struct {
	driver  string
	conexao string
	Db      *sql.DB
	erro    error
}

//Conexao estrutura com as propriedades de conexao ao Mysql
type Conexao repoMysql

//NovaConexao retorna o objeto que conecta ao banco de dados
func NovaConexao(driver string, conexao string) *Conexao {
	novoDb, e := sql.Open(driver, conexao)
	return &Conexao{
		driver:  driver,
		conexao: conexao,
		Db:      novoDb,
		erro:    e,
	}
}
