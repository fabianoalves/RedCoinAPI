package repositorio

import e "github.com/rteles86/RedCoinApi/API/entidade"

//TodosPerfil retorna todos os registros da tabela Perfil
func TodosPerfil(cn *Conexao) (listaPerfil []e.Perfil, erro error) {
	p := []e.Perfil{}

	rows, err := cn.Db.Query("SELECT idPerfil, perfil, registroApagado FROM Perfil")
	defer rows.Close()
	if err != nil {
		return p, err
	}

	for rows.Next() {
		pReg := e.Perfil{}
		rows.Scan(&pReg.IDPerfil, &pReg.Perfil, &pReg.RegistroApagado)
		p = append(p, pReg)
	}

	return p, err
}

//IDPerfil retorna o registro de um perfil de acordo com o ID informado
func IDPerfil(cn *Conexao, id int8) (perfil e.Perfil, erro error) {
	p := e.Perfil{}

	rows, err := cn.Db.Query("SELECT idPerfil, perfil FROM Perfil WHERE idPerfil = ?", id)
	defer rows.Close()

	if err != nil {
		return p, err
	}

	for rows.Next() {
		rows.Scan(&p.IDPerfil, &p.Perfil, &p.RegistroApagado)
	}

	return p, nil
}

//AdicionarPerfil método para adicionar um novo registro de Perfil
func AdicionarPerfil(cn *Conexao, perfil e.Perfil) (erro error) {

	addPerfil, err := cn.Db.Prepare("INSERT INTO Perfil(perfil)VALUES(?)")
	if err != nil {
		return err
	}

	addPerfil.Exec(perfil.Perfil)

	return nil
}

//AlterarPerfil método para atualizar o registro de um Perfil de acordo com ID informado
func AlterarPerfil(cn *Conexao, perfil e.Perfil) (erro error) {

	addPerfil, err := cn.Db.Prepare("UPDATE Perfil SET perfil = ?, registroApagado = ? WHERE idPerfil = ?")
	if err != nil {
		return err
	}

	addPerfil.Exec(perfil.Perfil, perfil.RegistroApagado, perfil.IDPerfil)

	return nil
}
