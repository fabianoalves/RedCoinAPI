package servico

import repo "github.com/rteles86/RedCoinApi/redcoin/repositorio"

var cn *repo.Conexao

//New cria a instância de conexão do repositório
func New(conexao *repo.Conexao) {
	cn = conexao
}
