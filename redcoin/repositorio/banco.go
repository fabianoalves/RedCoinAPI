package repositorio

import (
	"database/sql"
	"redcoin/conexoes"

	_ "github.com/go-sql-driver/mysql"
)

//Conexao retorna o objeto que conecta ao banco de dados
func Conexao() (db *sql.DB, erro error) {
	cn  := conexoes.CredenciaisMysql()

	driver := cn.Driver
	usuario := cn.Usuario
	senha := cn.Senha
	nomeBanco := cn.NomeBanco
	configuracaoTempo := cn.ConfiguracaoTempo
	porta := cn.Porta

	db, err := sql.Open(driver, usuario+":"+senha+"@"+porta+"/"+nomeBanco+configuracaoTempo)

	return db, err
}
