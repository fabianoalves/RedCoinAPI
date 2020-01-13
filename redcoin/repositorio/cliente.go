package repositorio

import (
	e "github.com/rteles86/RedCoinApi/redcoin/entidade"
)

//AdicionarCliente método para criar o registro do cliente para acessar a Api
func AdicionarCliente(cn *Conexao, cliente e.Cliente) (erro error) {
	addPerfil, e := cn.Db.Prepare("INSERT INTO ClienteApi(usuario, senha)VALUES(?,?)")
	addPerfil.Exec(cliente.Usuario, cliente.Senha)

	return e
}

//Cliente retornar os dados de um cliente a partir do nome de usuario
func Cliente(cn *Conexao, usuario string) e.Cliente {
	c := e.Cliente{}
	cn.Db.QueryRow(
		"SELECT usuario, senha FROM ClienteApi WHERE usuario = ?",
		usuario).Scan(
		&c.Usuario, &c.Senha)

	return c
}
