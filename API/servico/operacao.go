package servico

import (
	"time"

	i "github.com/rteles86/RedCoinApi/API/configuracoes/infra"
	e "github.com/rteles86/RedCoinApi/API/entidade"
	rep "github.com/rteles86/RedCoinApi/API/repositorio"
)

//EmailUsuarioOperacao método responsavel por listar as operações de um determinado usuario através de seu Email
func EmailUsuarioOperacao(email string) (uOperacao e.UsuarioOperacao, erro error) {
	return rep.EmailUsuarioOperacao(cn, email)
}

//PeriodoOperacao método responsavel por listar as operações de uma determinada Data
func PeriodoOperacao(data time.Time) (operacao []e.Operacao, erro error) {
	return rep.PeriodoOperacao(cn, data)
}

//PersistirOperacao método responsavel por adicionar uma operaçao
func PersistirOperacao(operacao e.Operacao) (erro error) {

	if i.ExisteCotacao() {
		operacao.ValorMoeda = i.CotacaoEmCache()
	} else {
		operacao.ValorMoeda, erro = i.CotacaoBitCoin()
		if erro != nil {
			return erro
		}
		i.CotacaoGravarCache(operacao.ValorMoeda)
	}
	operacao.ValorMoeda = operacao.ValorMoeda * operacao.ValorBitCoin

	return rep.AdicionarOperacao(cn, operacao)
}
