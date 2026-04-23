package repository

import (
	model "api/src/Model"
	"database/sql"
	"fmt"
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
func (repositorio Publicacoes) BuscarPorId(publicacaoID uint64) (model.Publicacao, error) {
	linha := repositorio.db.QueryRow(`
		SELECT 
			p.id,
			p.titulo,
			p.conteudo,
			p.autor_id,
			p.curtidas,
			p.criadaEm,
			u.nome,
			u.nick
		FROM publicacoes p
		INNER JOIN usuarios u ON u.id = p.autor_id
		WHERE p.id = ?
	`, publicacaoID)

	var publicacao model.Publicacao

	if err := linha.Scan(
		&publicacao.ID,
		&publicacao.Titulo,
		&publicacao.Conteudo,
		&publicacao.AutorID,
		&publicacao.Curtidas,
		&publicacao.CriadaEm,
		&publicacao.AutorNome,
		&publicacao.AutorNick,
	); err != nil {
		return model.Publicacao{}, err
	}

	return publicacao, nil
}

// Buscar é a função responsável por buscar todas as publicações de um usuário específico.
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]model.Publicacao, error) {
	linhas, err := repositorio.db.Query(`
		SELECT DISTINCT 
			p.id,
			p.titulo,
			p.conteudo,
			p.autor_id,
			p.curtidas,
			p.criadaEm,
			COALESCE(u.nome, ''),
			COALESCE(u.nick, '')
		FROM publicacoes p
		INNER JOIN usuarios u ON u.id = p.autor_id
		LEFT JOIN seguidores s ON p.autor_id = s.usuario_id
		WHERE p.autor_id = ? OR s.seguidor_id = ?
		ORDER BY p.id DESC
	`, usuarioID, usuarioID)

	if err != nil {
		fmt.Println("ERRO SQL:", err)
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []model.Publicacao

	for linhas.Next() {
		var publicacao model.Publicacao

		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNome,
			&publicacao.AutorNick,
		); err != nil {
			fmt.Println("ERRO SCAN:", err)
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Deletar é a função responsável por deletar uma publicação específica pelo ID.
func (repositorio Publicacoes) Deletar(publicacaoId uint64) error {
	statement, err := repositorio.db.Prepare(`delete from publicacoes where id = ?`)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(publicacaoId); err != nil {
		return err
	}

	return nil
}

// BuscarTodasPublicacoesPorUsuario é a função responsável por buscar todas as publicações de um usuário específico.
func (repositorio Publicacoes) BuscarTodasPublicacoesPorUsuario(usuarioId uint64) ([]model.Publicacao, error) {
	linhas, err := repositorio.db.Query(`select p.*, u.nick from publicacoes p join usuarios u on u.id = p.autor_id
	where p.autor_id = ?`,
		usuarioId)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	var publicacoes []model.Publicacao

	for linhas.Next() {
		var publicacao model.Publicacao

		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// CurtirPublicacao é a função responsável por adicionar um like à publicação de um usuário.
func (repositorio Publicacoes) CurtirPublicacao(publicacaoId uint64) error {
	statement, err := repositorio.db.Prepare(`update publicacoes set curtidas = curtidas + 1 where id = ?`)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(publicacaoId); err != nil {
		return err
	}

	return nil
}

// DeslikePublicacao remove uma curtida de uma publicacao especifica
func (repositorio Publicacoes) DeslikePublicacao(publicacaoId uint64) error {
	statement, err := repositorio.db.Prepare(`update publicacoes set curtidas =
	CASE WHEN curtidas > 0 THEN curtidas - 1 ELSE 0 END where id = ?`)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(publicacaoId); err != nil {
		return err
	}

	return nil
}

// Atualizar é a função responsável por atualizar uma publicação específica pelo ID.
func (repositorio Publicacoes) Atualizar(publicacaoId uint64, publicacao model.Publicacao) error {
	statement, err := repositorio.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoId)

	if err != nil {
		return err
	}

	return nil
}
