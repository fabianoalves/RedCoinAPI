package repositorio

import e "github.com/rteles86/RedCoinApi/API/entidade"

//TodosTipoOperacao retorna todos os registros da tabela tipoOperacao
func TodosTipoOperacao(cn *Conexao) (listaTipoOperacao []e.TipoOperacao, erro error) {
	to := []e.TipoOperacao{}

	rows, err := cn.Db.Query("SELECT idTipoOperacao, Operacao, registroApagado FROM TipoOperacao")
	defer rows.Close()
	if err != nil {
		return to, err
	}

	for rows.Next() {
		pTo := e.TipoOperacao{}
		rows.Scan(&pTo.IDTipoOperacao, &pTo.Operacao, &pTo.RegistroApagado)
		to = append(to, pTo)
	}

	return to, err
}

//IDTipoOperacao retorna o registro de um TipoOperacao de acordo com o ID informado
func IDTipoOperacao(cn *Conexao, id int8) (tipoOperacao e.TipoOperacao, erro error) {
	to := e.TipoOperacao{}

	e := cn.Db.QueryRow("SELECT idTipoOperacao, operacao FROM TipoOperacao WHERE idTipoOperacao = ?", id).Scan(&to.IDTipoOperacao, &to.Operacao)

	return to, e
}

//AdicionarTipoOperacao método para adicionar um novo registro de TipoOperacao
func AdicionarTipoOperacao(cn *Conexao, tipoOperacao e.TipoOperacao) (erro error) {

	addTipoOperacao, err := cn.Db.Prepare("INSERT INTO TipoOperacao(operacao)VALUES(?)")
	if err != nil {
		return err
	}

	addTipoOperacao.Exec(tipoOperacao.Operacao)

	return nil
}

//AlterarTipoOperacao método para atualizar o registro de um TipoOperacao de acordo com ID informado
func AlterarTipoOperacao(cn *Conexao, tipoOperacao e.TipoOperacao) (erro error) {

	addTipoOperacao, err := cn.Db.Prepare("UPDATE TipoOperacao SET operacao = ?, registroApagado = ? WHERE idTipoOperacao = ?")
	if err != nil {
		return err
	}

	addTipoOperacao.Exec(tipoOperacao.Operacao, tipoOperacao.RegistroApagado, tipoOperacao.IDTipoOperacao)

	return nil
}
