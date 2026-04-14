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
	return 0, nil
}
