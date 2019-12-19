package repositorio

import "redcoin/modelos"

//TodosTipoOperacao retorna todos os registros da tabela tipoOperacao
func TodosTipoOperacao() (listaTipoOperacao []modelos.TipoOperacao, erro error) {
	to := []modelos.TipoOperacao{}
	db, err := Conexao()
	if err != nil {
		return to, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT idTipoOperacao, Operacao, registroApagado FROM TipoOperacao")
	defer rows.Close()
	if err != nil {
		return to, err
	}

	for rows.Next() {
		pTo := modelos.TipoOperacao{}
		rows.Scan(&pTo.IDTipoOperacao, &pTo.Operacao, &pTo.RegistroApagado)
		to = append(to, pTo)
	}

	return to, err
}

//IDTipoOperacao retorna o registro de um TipoOperacao de acordo com o ID informado
func IDTipoOperacao(id int8) (tipoOperacao modelos.TipoOperacao, erro error) {
	to := modelos.TipoOperacao{}

	db, err := Conexao()
	if err != nil {
		return to, err
	}
	defer db.Close()

	db.QueryRow("SELECT idTipoOperacao, operacao FROM TipoOperacao WHERE idTipoOperacao = ?", id).Scan(&to.IDTipoOperacao, &to.Operacao)

	return to, err
}

//AdicionarTipoOperacao método para adicionar um novo registro de TipoOperacao
func AdicionarTipoOperacao(tipoOperacao modelos.TipoOperacao) (erro error) {
	db, err := Conexao()
	if err != nil {
		return err
	}
	defer db.Close()

	addTipoOperacao, err := db.Prepare("INSERT INTO TipoOperacao(operacao)VALUES(?)")
	if err != nil {
		return err
	}

	addTipoOperacao.Exec(tipoOperacao.Operacao)

	return nil
}

//AlterarTipoOperacao método para atualizar o registro de um TipoOperacao de acordo com ID informado
func AlterarTipoOperacao(tipoOperacao modelos.TipoOperacao) (erro error) {
	db, err := Conexao()
	if err != nil {
		return err
	}
	defer db.Close()

	addTipoOperacao, err := db.Prepare("UPDATE TipoOperacao SET operacao = ?, registroApagado = ? WHERE idTipoOperacao = ?")
	if err != nil {
		return err
	}

	addTipoOperacao.Exec(tipoOperacao.Operacao, tipoOperacao.RegistroApagado, tipoOperacao.IDTipoOperacao)

	return nil
}
