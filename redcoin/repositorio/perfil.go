package repositorio

import "redcoin/modelos"

//TodosPerfil retorna todos os registros da tabela Perfil
func TodosPerfil() (listaPerfil []modelos.Perfil, erro error) {
	p := []modelos.Perfil{}
	db, err := Conexao()
	if err != nil {
		return p, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT idPerfil, perfil, registroApagado FROM Perfil")
	defer rows.Close()
	if err != nil {
		return p, err
	}

	for rows.Next() {
		pReg := modelos.Perfil{}
		rows.Scan(&pReg.IdPerfil, &pReg.Perfil, &pReg.RegistroApagado)
		p = append(p, pReg)
	}

	return p, err
}

//IDPerfil retorna o registro de um perfil de acordo com o ID informado
func IDPerfil(id int8) (perfil modelos.Perfil, erro error) {
	p := modelos.Perfil{}

	db, err := Conexao()
	if err != nil {
		return p, err
	}
	defer db.Close()

	db.QueryRow("SELECT idPerfil, perfil FROM Perfil WHERE idPerfil = ?", id).Scan(&p.IdPerfil, &p.Perfil)

	return p, err
}

//AdicionarPerfil método para adicionar um novo registro de Perfil
func AdicionarPerfil(perfil modelos.Perfil) (erro error) {
	db, err := Conexao()
	if err != nil {
		return err
	}
	defer db.Close()

	addPerfil, err := db.Prepare("INSERT INTO Perfil(perfil)VALUES(?)")
	if err != nil {
		return err
	}

	addPerfil.Exec(perfil.Perfil)

	return nil
}

//AlterarPerfil método para atualizar o registro de um Perfil de acordo com ID informado
func AlterarPerfil(perfil modelos.Perfil) (erro error) {
	db, err := Conexao()
	if err != nil {
		return err
	}
	defer db.Close()

	addPerfil, err := db.Prepare("UPDATE Perfil SET perfil = ?, registroApagado = ? WHERE idPerfil = ?")
	if err != nil {
		return err
	}

	addPerfil.Exec(perfil.Perfil, perfil.RegistroApagado, perfil.IdPerfil)

	return nil
}
