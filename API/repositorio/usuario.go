package repositorio

import e "github.com/rteles86/RedCoinApi/API/entidade"

//TodosUsuario retorna todos os registros da tabela Usuari0
func TodosUsuario(cn *Conexao) (listaUsuario []e.Usuario, erro error) {
	u := []e.Usuario{}

	rows, err := cn.Db.Query(`
	SELECT
	u.idUsuario
    ,u.email
    ,u.senha
    ,u.nome	
    ,u.ultimoNome
    ,u.dataNascimento
    ,u.quantidadeMoeda
    ,p.idPerfil
    ,p.perfil
		FROM
	Usuario						AS	u
    INNER JOIN	PerfilUsuario	AS	pu
		ON	pu.idUsuario 	= u.idUsuario
    INNER JOIN	Perfil			AS	p
		ON	p.idPerfil 		= pu.idPerfil
	ORDER BY
		u.IDUsuario ASC
	`)
	defer rows.Close()
	if err != nil {
		return u, err
	}

	var indexU, idUsuario int = -1, 0
	for rows.Next() {
		pU := e.Usuario{}
		pP := e.Perfil{}
		rows.Scan(&pU.IDUsuario, &pU.Email, &pU.Senha, &pU.Nome, &pU.UltimoNome, &pU.DataNascimento,
			&pU.QuantidadeMoeda, &pP.IDPerfil, &pP.Perfil)

		if idUsuario != pU.IDUsuario {
			u = append(u, pU)
			indexU++
			idUsuario = pU.IDUsuario
			u[indexU].PerfilUsuario = append(u[indexU].PerfilUsuario, pP)
		} else {
			u[indexU].PerfilUsuario = append(u[indexU].PerfilUsuario, pP)
		}
	}

	return u, err
}

//IDUsuario retorna o registro de um usuario de acord com o ID informado
func IDUsuario(cn *Conexao, id int) (usuario e.Usuario, erro error) {
	u := e.Usuario{}
	p := e.Perfil{}

	rows, err := cn.Db.Query(`
	SELECT
	u.idUsuario
    ,u.email
    ,u.senha
    ,u.nome	
    ,u.ultimoNome
    ,u.dataNascimento
    ,u.quantidadeMoeda
    ,p.idPerfil
    ,p.perfil
		FROM
	Usuario						AS	u
    INNER JOIN	PerfilUsuario	AS	pu
		ON	pu.idUsuario 	= u.idUsuario
    INNER JOIN	Perfil			AS	p
		ON	p.idPerfil 		= pu.idPerfil
	WHERE
		u.idUsuario = ?
	`, id)

	for rows.Next() {
		rows.Scan(&u.IDUsuario, &u.Email, &u.Senha, u.Nome, &u.UltimoNome, &u.DataNascimento, &u.QuantidadeMoeda,
			&p.IDPerfil, &p.Perfil)
		u.PerfilUsuario = append(u.PerfilUsuario, p)
	}

	return u, err
}

//AdicionarUsuario método para adicionar um novo regitro de Usuario
func AdicionarUsuario(cn *Conexao, usuario e.Usuario) (erro error) {

	addUsuario, err := cn.Db.Prepare("INSERT INTO Usuario(email, senha, nome, ultimoNome, dataNascimento, quantidadeMoeda)VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	uDb, err := addUsuario.Exec(usuario.Email, usuario.Senha, usuario.Nome, usuario.UltimoNome, usuario.DataNascimento,
		usuario.QuantidadeMoeda)
	if err != nil {
		return err
	}

	idUsuario, _ := uDb.LastInsertId()

	for _, p := range usuario.PerfilUsuario {
		var pu e.PerfilUsuario
		pu.IDPerfil = p.IDPerfil
		pu.IDUsuario = int(idUsuario)

		err := AdicionarPerfilUsuario(cn, pu)
		if err != nil {
			return err
		}
	}

	return nil
}

//AlterarUsuario método para atualizar o registro do Usuario
func AlterarUsuario(cn *Conexao, usuario e.Usuario) (erro error) {

	addUsuario, err := cn.Db.Prepare(`UPDATE Usuario SET 
		email = ?, senha = ?, nome = ?, ultimoNome = ?, dataNascimento = ?, quantidadeMoeda = ? 
		WHERE idUsuario = ?`)
	if err != nil {
		return err
	}

	addUsuario.Exec(usuario.Email, usuario.Senha, usuario.Nome, usuario.UltimoNome, usuario.DataNascimento,
		usuario.QuantidadeMoeda, usuario.IDUsuario)

	for _, p := range usuario.PerfilUsuario {
		var pu e.PerfilUsuario
		pu.IDPerfil = p.IDPerfil
		pu.IDUsuario = usuario.IDUsuario

		err := AdicionarPerfilUsuario(cn, pu)
		if err != nil {
			return err
		}
	}
	return nil
}

func saldoBitCoin(cn *Conexao, usuario e.Usuario) (erro error) {

	addUsuario, err := cn.Db.Prepare(`UPDATE Usuario SET quantidadeMoeda = quantidadeMoeda + ? WHERE idUsuario = ?`)
	if err != nil {
		return err
	}

	_, err = addUsuario.Exec(usuario.QuantidadeMoeda, usuario.IDUsuario)

	return err
}

//AtualizaSaldoBitCoin Atualiza as quantidades de BitCoin do Comprador e do Vend edor após Op eacao
func AtualizaSaldoBitCoin(cn *Conexao, valorBitCoin float64, idVendedor int, idComprador int) (erro error) {
	var comprador e.Usuario
	comprador.IDUsuario = idComprador
	comprador.QuantidadeMoeda = valorBitCoin

	var vendedor e.Usuario
	vendedor.IDUsuario = idVendedor
	vendedor.QuantidadeMoeda = -valorBitCoin

	err := saldoBitCoin(cn, comprador)
	if err != nil {
		return err
	}

	return saldoBitCoin(cn, vendedor)
}
