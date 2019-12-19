package repositorio

import (
	"redcoin/modelos"
	"time"
)

//Usuario estrutura qe representa a tabela Usuario
type Usuario struct {
	IdUsuario       int
	Email           string
	Senha           string
	Nome            string
	UltimoNome      string
	DataNascimento  time.Time
	QuantidadeMoeda float64
	RegistroApagado bool
	PerfilUsuario   []modelos.Perfil
}

//TodosUsuario retorna todos os registros da tabela Usuari
func TodosUsuario() (listaUsuario []modelos.Usuario, erro error) {
	u := []modelos.Usuario{}
	db, err := Conexao()
	if err != nil {
		return u, err
	}
	defer db.Close()

	rows, err := db.Query(`
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
		u.IdUsuario ASC
	`)
	defer rows.Close()
	if err != nil {
		return u, err
	}

	var indexU, idUsuario int = -1, 0
	for rows.Next() {
		pU := modelos.Usuario{}
		pP := modelos.Perfil{}
		rows.Scan(&pU.IdUsuario, &pU.Email, &pU.Senha, &pU.Nome, &pU.UltimoNome, &pU.DataNascimento,
			&pU.QuantidadeMoeda, &pP.IdPerfil, &pP.Perfil)

		if idUsuario != pU.IdUsuario {
			u = append(u, pU)
			indexU += 1
			idUsuario = pU.IdUsuario
			u[indexU].PerfilUsuario = append(u[indexU].PerfilUsuario, pP)
		} else {
			u[indexU].PerfilUsuario = append(u[indexU].PerfilUsuario, pP)
		}
	}

	return u, err
}

//IDUsuario retorna o registro de um usuario de acord com o ID informado
func IDUsuario(id int) (usuario modelos.Usuario, erro error) {
	u := modelos.Usuario{}
	p := modelos.Perfil{}

	db, err := Conexao()
	if err != nil {
		return u, err
	}
	defer db.Close()

	rows, err := db.Query(`
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
		rows.Scan(&u.IdUsuario, &u.Email, &u.Senha, u.Nome, &u.UltimoNome, &u.DataNascimento, &u.QuantidadeMoeda,
			&p.IdPerfil, &p.Perfil)
		u.PerfilUsuario = append(u.PerfilUsuario, p)
	}

	return u, err
}

//AdicionarUsuario método para adicionar um novo regitro de Usuario
func AdicionarUsuario(usuario modelos.Usuario) (erro error) {
	db, err := Conexao()
	if err != nil {
		return err
	}
	defer db.Close()

	addUsuario, err := db.Prepare("INSERT INTO Usuario(email, senha, nome, ultimoNome, dataNascimento, quantidadeMoeda)VALUES(?, ?, ?, ?, ?, ?)")
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
		var pu modelos.PerfilUsuario
		pu.IdPerfil = p.IdPerfil
		pu.IdUsuario = int(idUsuario)

		err := AdicionarPerfilUsuario(pu)
		if err != nil {
			return err
		}
	}

	return nil
}

//AlterarUsuario método para atualizar o registro do Usuario
func AlterarUsuario(usuario modelos.Usuario) (erro error) {
	db, err := Conexao()
	if err != nil {
		return err
	}
	defer db.Close()

	addUsuario, err := db.Prepare(`UPDATE Usuario SET 
		email = ?, senha = ?, nome = ?, ultimoNome = ?, dataNascimento = ?, quantidadeMoeda = ? 
		WHERE idUsuario = ?`)
	if err != nil {
		return err
	}

	addUsuario.Exec(usuario.Email, usuario.Senha, usuario.Nome, usuario.UltimoNome, usuario.DataNascimento,
		usuario.QuantidadeMoeda, usuario.IdUsuario)

	for _, p := range usuario.PerfilUsuario {
		var pu modelos.PerfilUsuario
		pu.IdPerfil = p.IdPerfil
		pu.IdUsuario = usuario.IdUsuario

		err := AdicionarPerfilUsuario(pu)
		if err != nil {
			return err
		}
	}
	return nil
}

func saldoBitCoin(usuario modelos.Usuario) (erro error) {
	db, err := Conexao()
	if err != nil {
		return err
	}
	defer db.Close()

	addUsuario, err := db.Prepare(`UPDATE Usuario SET quantidadeMoeda = quantidadeMoeda + ? WHERE idUsuario = ?`)
	if err != nil {
		return err
	}

	_, err = addUsuario.Exec(usuario.QuantidadeMoeda, usuario.IdUsuario)

	return err
}

//AtualizaSaldoBitCoin Atualiza as quantidades de BitCoin do Comprador e do Vend edor após Op eacao
func AtualizaSaldoBitCoin(valorBitCoin float64, idVendedor int, idComprador int) (erro error) {
	var comprador modelos.Usuario
	comprador.IdUsuario = idComprador
	comprador.QuantidadeMoeda = valorBitCoin

	var vendedor modelos.Usuario
	vendedor.IdUsuario = idVendedor
	vendedor.QuantidadeMoeda = -valorBitCoin

	err := saldoBitCoin(comprador)
	if err != nil {
		return err
	}

	return saldoBitCoin(vendedor)
}
