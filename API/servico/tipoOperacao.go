package servico

import (
	e "github.com/rteles86/RedCoinApi/redcoin/entidade"
	repo "github.com/rteles86/RedCoinApi/redcoin/repositorio"
)

//TodosTipoOperacao retorna a lista com todos dados de TipoOperacao
func TodosTipoOperacao() (listaTipoOperacao []e.TipoOperacao, erro error) {
	listaTipoOperacao, erro = repo.TodosTipoOperacao(cn)

	return listaTipoOperacao, erro
}

//IDTipoOperacao retorna o TipoOperacao de acordo com o ID informado
func IDTipoOperacao(id int8) (tipoOperacao e.TipoOperacao, erro error) {
	return repo.IDTipoOperacao(cn, id)

}

//AdicionarTipoOperacao solicita a criação de um registro de TipoOperacao
func AdicionarTipoOperacao(tipoOperacao e.TipoOperacao) (erro error) {
	return repo.AdicionarTipoOperacao(cn, tipoOperacao)
}

//AtualizarTipoOperacao solicita a alteração de um registro de TipoOperacao
func AtualizarTipoOperacao(tipoOperacao e.TipoOperacao) (erro error) {
	return repo.AlterarTipoOperacao(cn, tipoOperacao)
}
