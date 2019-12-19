package repositorio

import "redcoin/modelos"

//IDUsuarioPerfilUsurio retorna todos os registros da tabela PerfilUsuario de acordo com o idUsuario
func IDUsuarioPerfilUuario(idusuario int) (perfilUsuario []modelos.PerfilUsuario, erro error) {
	pu := []modelos.PerfilUsuario{}
	db, err := Conexao()
	if err != nil {
		return pu, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT idPerfilUsuario, idPerfil, idUsuario FROM PerfilUsuario WHERE idUsuario = ?", idusuario)
	defer rows.Close()
	if err != nil {
		return pu, err
	}

	for rows.Next() {
		perfilU := modelos.PerfilUsuario{}
		rows.Scan(&perfilU.IdPerfilUsuario, &perfilU.IdPerfil, &perfilU.IdUsuario)
		pu = append(pu, perfilU)
	}

	return pu, err
}

//AdicionarPerfilUsuario método para adicionar um novo registro de PerfilUsuario
func AdicionarPerfilUsuario(perfilUsuario modelos.PerfilUsuario) (erro error) {
	db, err := Conexao()
	if err != nil {
		return err
	}
	defer db.Close()

	addPerfilUsuario, err := db.Prepare(`
	INSERT INTO PerfilUsuario(idPerfil, idUsuario)VALUES(?, ?)
	ON DUPLICATE KEY UPDATE idPerfil = ?, idUsuario = ?
	`)
	if err != nil {
		return err
	}

	addPerfilUsuario.Exec(perfilUsuario.IdPerfil, perfilUsuario.IdUsuario, perfilUsuario.IdPerfil, perfilUsuario.IdUsuario)

	return nil
}
