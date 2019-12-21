package repositorio

import (
	"redcoin/modelos"
)

//AdicionarClienteApi método para criar o acesso do client Api
func AdicionarClienteApi(clienteApi modelos.ClienteApi) (erro error) {
	db, err := Conexao()
	if err != nil {
		return err
	}
	defer db.Close()

	addPerfil, err := db.Prepare("INSERT INTO ClienteApi(usuario, senha)VALUES(?,?)")
	if err != nil {
		return err
	}

	addPerfil.Exec(clienteApi.Usuario, clienteApi.Senha)

	return nil
}

//AutenticarClienteApi - métod para verificar as credenciais do cliente da api
func AutenticarClienteApi(usuario string, senha string) (existe bool, erro error) {
	ca := modelos.ClienteApi{}

	db, err := Conexao()
	if err != nil {
		return false, err
	}
	defer db.Close()

	db.QueryRow("SELECT usuario FROM ClienteApi WHERE usuario = ? AND senha = ?", usuario, senha).Scan(&ca.Usuario)

	return bool(ca.Usuario == usuario), err
}
