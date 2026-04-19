package repository

import (
	model "api/src/Model"
	"database/sql"
)

// RepositorioPublicacoes representa um repositorio de publicações
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioPublicacoes cria um novo repositorio de publicações
func NovoRepositorioPublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar é a função responsável por criar uma nova publicação no banco de dados.
func (repositorio Publicacoes) Criar(publicacao model.Publicacao) (uint64, error) {
	statement, err := repositorio.db.Prepare("INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)

	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil

}
