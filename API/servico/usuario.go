package servico

import (
	"errors"

	u "github.com/rteles86/RedCoinApi/redcoin/configuracoes/utils"
	e "github.com/rteles86/RedCoinApi/redcoin/entidade"
	repo "github.com/rteles86/RedCoinApi/redcoin/repositorio"
)

//TodosUsuario retorna a lista com todos dados de Usuario
func TodosUsuario() (listaUsuario []e.Usuario, erro error) {
	listaUsuario, erro = repo.TodosUsuario(cn)

	return listaUsuario, erro
}

//IDUsuario retorna o Usuario de acordo com o ID informado
func IDUsuario(id int) (usuario e.Usuario, erro error) {
	if id <= 0 {
		return usuario, errors.New("O ID do usuario deve ser maior que ZERO")
	}
	return repo.IDUsuario(cn, id)

}

//AdicionarUsuario solicita a criação de um registro de Usuario
func AdicionarUsuario(usuario e.Usuario) (erro error) {
	usuario.Senha = u.CriptografarSenha(usuario.Senha)
	return repo.AdicionarUsuario(cn, usuario)
}

//AtualizarUsuario solicita a alteração de um registro de Usuario
func AtualizarUsuario(usuario e.Usuario) (erro error) {
	usuario.Senha = u.CriptografarSenha(usuario.Senha)
	return repo.AlterarUsuario(cn, usuario)
}
