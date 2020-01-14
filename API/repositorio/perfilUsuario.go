package repositorio

import e "github.com/rteles86/RedCoinApi/API/entidade"

//IDUsuarioPerfilUsuario retorna todos os registros da tabela PerfilUsuario de acordo com o idUsuario
func IDUsuarioPerfilUsuario(cn *Conexao, idusuario int) (perfilUsuario []e.PerfilUsuario, erro error) {
	pu := []e.PerfilUsuario{}

	rows, err := cn.Db.Query("SELECT idPerfilUsuario, idPerfil, idUsuario FROM PerfilUsuario WHERE idUsuario = ?", idusuario)
	defer rows.Close()
	if err != nil {
		return pu, err
	}

	for rows.Next() {
		perfilU := e.PerfilUsuario{}
		rows.Scan(&perfilU.IDPerfilUsuario, &perfilU.IDPerfil, &perfilU.IDUsuario)
		pu = append(pu, perfilU)
	}

	return pu, err
}

//AdicionarPerfilUsuario método para adicionar um novo registro de PerfilUsuario
func AdicionarPerfilUsuario(cn *Conexao, perfilUsuario e.PerfilUsuario) (erro error) {

	addPerfilUsuario, err := cn.Db.Prepare(`
	INSERT INTO PerfilUsuario(idPerfil, idUsuario)VALUES(?, ?)
	ON DUPLICATE KEY UPDATE idPerfil = ?, idUsuario = ?
	`)
	if err != nil {
		return err
	}

	addPerfilUsuario.Exec(perfilUsuario.IDPerfil, perfilUsuario.IDUsuario, perfilUsuario.IDPerfil, perfilUsuario.IDUsuario)

	return nil
}
