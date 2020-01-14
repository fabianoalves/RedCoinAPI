package repositorio

import (
	"time"

	e "github.com/rteles86/RedCoinApi/redcoin/entidade"
)

//AdicionarOperacao método para adicionar um novo registro de Operacao
func AdicionarOperacao(cn *Conexao, operacao e.Operacao) (erro error) {

	addOperacao, err := cn.Db.Prepare(`INSERT INTO Operacao
										(idTipoOperacao
										, idVendedor
										, idComprador
										, dataOperacao
										, valorMoeda
										, valorBitCoin)VALUES(?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	_, err = addOperacao.Exec(operacao.TipoOperacao.IDTipoOperacao, operacao.Vendedor.IDUsuario, operacao.Comprador.IDUsuario, operacao.DataOperacao, operacao.ValorMoeda, operacao.ValorBitCoin)
	if err != nil {
		return err
	}

	err = AtualizaSaldoBitCoin(cn, operacao.ValorBitCoin, operacao.Vendedor.IDUsuario, operacao.Comprador.IDUsuario)
	if err != nil {
		return err
	}

	return nil
}

//EmailUsuarioOperacao retorna todos os registros de operacao realizado por determinado email de um usuario
func EmailUsuarioOperacao(cn *Conexao, email string) (usuarioOperacao e.UsuarioOperacao, erro error) {
	uO := e.UsuarioOperacao{}

	rows, err := cn.Db.Query(`
	 SELECT
		u.idUsuario 
		,CONCAT(u.nome, ' ', u.ultimoNome) AS nomeCompleto
		,o.idOperacao 
		,t.operacao	AS	tipoOperacao 
		,v.idUsuario AS idVendendor 
		,v.email	AS emailVendedor 
		,CONCAT(v.nome, ' ', v.ultimoNome) AS nomeCompletoVendedor 
		,o.dataOperacao
		,o.valorMoeda 
		,o.valorBitCoin 
	FROM 
		Usuario AS u  
		INNER JOIN	Operacao AS o  
			ON	o.idComprador = u.idUsuario 
		INNER JOIN	TipoOperacao	AS t  
			ON	t.idTipoOperacao = o.idTipoOperacao 
		INNER JOIN	Usuario AS v  
			ON	v.idUsuario = o.idVendedor 
	WHERE 
		u.email = ?;`, email)
	defer rows.Close()
	if err != nil {
		return uO, err
	}

	for rows.Next() {
		uoRow := e.Operacoes{}
		rows.Scan(&uO.IDUsuario, &uO.NomeCompleto, &uoRow.IDOperacao, &uoRow.TipoOperacao, &uoRow.IDVendedor,
			&uoRow.EmailVendedor, &uoRow.NomeCompletoVendedor, &uoRow.DataOperacao, &uoRow.ValorMoeda, &uoRow.ValorBitCoin)
		uO.Email = email
		uO.Operacoes = append(uO.Operacoes, uoRow)
	}

	return uO, err
}

//PeriodoOperacao retorna as operações  de acordo com um periodo informad o
func PeriodoOperacao(cn *Conexao, periodo time.Time) (operacao []e.Operacao, erro error) {
	o := []e.Operacao{}
	anoI, mesI, diaI := periodo.Date()

	inicio := time.Date(anoI, mesI, diaI, 0, 0, 0, 0, time.UTC)
	fim := time.Date(anoI, mesI, diaI, 23, 59, 59, 997, time.UTC)

	rows, err := cn.Db.Query(`select  
			o.idOperacao
			,t.idTipoOperacao
			,t.operacao
			,v.email AS emailVendedor
			,v.nome AS nomeVendedor 
			,v.ultimoNome AS ultimoNomeVendedor
			,c.email AS emailComprador
			,c.nome AS nomeComprador
			, c.ultimoNome AS ultimoNomeComprador
			,o.dataOperacao
			,o.valorMoeda
			,o.valorBitCoin
		FROM Operacao AS o
			INNER JOIN	TipoOperacao AS t
				ON	t.idTipoOperacao = o.idTipoOperacao
			INNER JOIN	Usuario AS v
				ON	v.idUsuario = o.idVendedor
			INNER JOIN	Usuario AS c
				ON	c.idUsuario = o.idComprador
		WHERE
		o.dataOperacao	BETWEEN ? AND ?`, inicio, fim)

	defer rows.Close()
	if err != nil {
		return o, err
	}

	for rows.Next() {
		oRow := e.Operacao{}

		rows.Scan(&oRow.IDOperacao, &oRow.TipoOperacao.IDTipoOperacao, &oRow.TipoOperacao.Operacao,
			&oRow.Vendedor.Email, &oRow.Vendedor.Nome, &oRow.Vendedor.UltimoNome,
			&oRow.Comprador.Email, &oRow.Comprador.Nome, &oRow.Comprador.UltimoNome,
			&oRow.DataOperacao, &oRow.ValorMoeda, &oRow.ValorBitCoin)

		o = append(o, oRow)
	}

	return o, err
}
