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

// BuscarPorId é a função responsável por buscar uma publicação específica pelo ID.
func (repositorio Publicacoes) BuscarPorId(PublicacoesId uint64) (model.Publicacao, error) {
	linha, err := repositorio.db.Query(`
		select p.*, u.nick from publicacoes p inner join usuarios u on u.id = p.autor_id where p.id= ?
	`, PublicacoesId)

	if err != nil {
		return model.Publicacao{}, err
	}

	defer linha.Close()

	var publicacao model.Publicacao

	if linha.Next() {
		if err = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return model.Publicacao{}, err
		}
	}

	return publicacao, nil
}
