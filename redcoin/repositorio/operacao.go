package repositorio

import (
	_ "fmt"
	"time"
	"redcoin/modelos"
)

//AdicionarOperacao método para adicionar um novo registro de Operacao
func AdicionarOperacao(operacao modelos.Operacao) (erro error) {
	db, err := Conexao()
	if err != nil {
		return err
	}
	defer db.Close()

	addOperacao, err := db.Prepare(`INSERT INTO Operacao
										(idTipoOperacao
										, idVendedor
										, idComprador
										, dataOperacao
										, valorMoeda
										, valorBitCoin)VALUES(?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	_ ,err = addOperacao.Exec(operacao.TipoOperacao.IDTipoOperacao, operacao.Vendedor.IdUsuario, operacao.Comprador.IdUsuario, operacao.DataOperacao, operacao.ValorMoeda, operacao.ValorBitCoin)
	if err != nil{
		return err
	}

	err = AtualizaSaldoBitCoin(operacao.ValorBitCoin, operacao.Vendedor.IdUsuario, operacao.Comprador.IdUsuario)
	if err != nil{
		return err
	}

	return nil
}

//EmailUsuarioOperacao retorna todos os registros de operacao realizado por determinado email de um usuario
func EmailUsuarioOperacao(email string) (usuarioOperacao modelos.UsuarioOperacao, erro error) {
	uO := modelos.UsuarioOperacao{}

	db, err := Conexao()
	if err != nil {
		return uO, err
	}
	defer db.Close()

	rows, err := db.Query(`
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
		uoRow := modelos.Operacoes{}
		rows.Scan(&uO.IdUsuario, &uO.NomeCompleto, &uoRow.IdOperacao, &uoRow.TipoOperacao, &uoRow.IdVendedor, 
			&uoRow.EmailVendedor, &uoRow.NomeCompletoVendedor, &uoRow.DataOperacao, &uoRow.ValorMoeda, &uoRow.ValorBitCoin)
		uO.Email = email
		uO.Operacoes = append(uO.Operacoes, uoRow)
	}

	return uO, err
}

//PeriodoOperacao retorna as operações  de acordo com um periodo informad o
func PeriodoOperacao(periodo time.Time)  (operacao []modelos.Operacao, erro error ) {
	o := []modelos.Operacao{}
	anoI, mesI, diaI := periodo.Date()

	inicio := time.Date(anoI, mesI, diaI, 0, 0, 0, 0, time.UTC)
	fim := time.Date(anoI, mesI, diaI, 23, 59, 59, 997, time.UTC)

	db, err := Conexao()
	if err != nil {
		return o, err
	}
	defer db.Close()

	rows, err := db.Query(`select  
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
		oRow := modelos.Operacao{}

		rows.Scan(&oRow.IdOperacao, &oRow.TipoOperacao.IDTipoOperacao, &oRow.TipoOperacao.Operacao,
			&oRow.Vendedor.Email, &oRow.Vendedor.Nome, &oRow.Vendedor.UltimoNome,
			&oRow.Comprador.Email, &oRow.Comprador.Nome, &oRow.Comprador.UltimoNome,
			&oRow.DataOperacao, &oRow.ValorMoeda, &oRow.ValorBitCoin)

		o = append(o, oRow)
	}

	return o, err
}