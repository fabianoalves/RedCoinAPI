package servico

import (
	e "github.com/rteles86/RedCoinApi/API/entidade"
	repo "github.com/rteles86/RedCoinApi/API/repositorio"
)

//TodosPerfil retorna a lista com todos dados de Perfil
func TodosPerfil() (listaPerfil []e.Perfil, erro error) {
	listaPerfil, erro = repo.TodosPerfil(cn)

	return listaPerfil, erro
}

//IDPerfil retorna o perfil de acordo com o ID informado
func IDPerfil(id int8) (perfil e.Perfil, erro error) {
	return repo.IDPerfil(cn, id)
}

//AdicionarPerfil solicita a criação de um registro de perfil
func AdicionarPerfil(perfil e.Perfil) (erro error) {
	return repo.AdicionarPerfil(cn, perfil)
}

//AtualizarPerfil solicita a alteração de um registro de perfil
func AtualizarPerfil(perfil e.Perfil) (erro error) {
	return repo.AlterarPerfil(cn, perfil)
}
