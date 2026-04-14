package repository

import (
	model "api/src/Model"
	"database/sql"
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
