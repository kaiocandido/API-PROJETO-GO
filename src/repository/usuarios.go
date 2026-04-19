package repository

import (
	model "api/src/Model"
	"database/sql"
	"fmt"
)

// Usuarios representa um repositorio
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioUsuarios cria um repositorio de usuarios
func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuario no banco de dados
func (u Usuarios) Criar(usuario model.Usuario) (uint64, error) {
	statement, err := u.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resul, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resul.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil

}

// Buscar traz todos usuarios de acordo com o filtro
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]model.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, err := repositorio.db.Query(
		"select id, nome, nick, email, criado from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var usuarios []model.Usuario

	for linhas.Next() {
		var usuario model.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Criado,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarPorId traz um usuario de acordo com o id
func (repositorio Usuarios) BuscarPorId(IdUsuario uint64) (model.Usuario, error) {

	linhas, err := repositorio.db.Query(
		"select  id, nome, nick, email, criado from usuarios where id = ?",
		IdUsuario,
	)

	if err != nil {
		return model.Usuario{}, err
	}

	defer linhas.Close()

	var usuario model.Usuario

	if linhas.Next() {
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Criado,
		); err != nil {
			return model.Usuario{}, err
		}
	}

	return usuario, nil

}

// Atualizar serve para alterar um usuario atraves do ID
func (repositorio Usuarios) Atualizar(ID uint64, usuario model.Usuario) error {
	statement, err := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}

	return nil
}

// Deletar exclui um usuario do banco
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare(
		"delete from usuarios where id = ?",
	)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// BuscarPorEmail tras um usuario através do email
func (repositorio Usuarios) BuscarPorEmail(email string) (model.Usuario, error) {
	linha, err := repositorio.db.Query("select id, senha from usuarios where email = ?", email)

	if err != nil {
		return model.Usuario{}, err
	}

	defer linha.Close()

	var usuario model.Usuario

	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return model.Usuario{}, err
		}
	}

	return usuario, nil
}

// Seguir permite que um usuario siga outro
func (repositorio Usuarios) Seguir(usuarioId, seguidorId uint64) error {
	statement, err := repositorio.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(usuarioId, seguidorId); err != nil {
		return err
	}

	return nil
}

// PararDeSeguirUsuario permite que um usuario deixe de seguir outro usuario
func (repositorio Usuarios) PararDeSeguirUsuario(usuarioId, seguidorId uint64) error {
	statement, err := repositorio.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(usuarioId, seguidorId); err != nil {
		return err
	}

	return nil

}

// BuscarSeguidores traz os seguidores de um usuario
func (repositorio Usuarios) BuscarSeguidores(usuarioId uint64) ([]model.Usuario, error) {
	linhas, err := repositorio.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.criado from usuarios u
	inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?`,
		usuarioId,
	)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var usuarios []model.Usuario

	for linhas.Next() {
		var usuario model.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Criado,
		); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarSeguindo traz os usuarios que um usuario segue
func (repositorio Usuarios) BuscarSeguindo(usuarioId uint64) ([]model.Usuario, error) {
	linhas, err := repositorio.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.criado from usuarios u
	inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?`,
		usuarioId,
	)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var usuarios []model.Usuario

	for linhas.Next() {
		var usuario model.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Criado,
		); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarSenha tras a senha de um usuario
func (repositorio Usuarios) BuscarSenha(usuarioId uint64) (string, error) {
	linha, err := repositorio.db.Query("select senha from usuarios where id = ?")

	if err != nil {
		return "", err
	}

	defer linha.Close()

	var usuario model.Usuario

	if linha.Next() {
		if err = linha.Scan(&usuario.Senha); err != nil {
			return "", err
		}
	}

	return usuario.Senha, nil
}

// AlterarSenha altera a senha de um usuario
func (repositorio Usuarios) AlterarSenha(usuarioId uint64, senha string) error {
	statement, err := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(senha, usuarioId); err != nil {
		return err
	}

	return nil
}
