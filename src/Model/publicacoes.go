package model

import (
	"errors"
	"strings"
	"time"
)

// Publicacao representa uma publicação feita por um usuário.
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	Curtidas  uint64    `json:"curtidas,omitempty"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	AutorNome string    `json:"autorNome,omitempty"`
}

// Preparar é a função responsável por preparar os dados da publicação, validando e formatando os campos antes de serem armazenados no banco de dados.
func (publicacao *Publicacao) Preparar() error {
	if err := publicacao.validar(); err != nil {
		return err
	}

	publicacao.formatar()

	return nil
}

// validar é a função responsável por validar os campos da publicação, garantindo que eles estejam preenchidos corretamente.
func (publicacao *Publicacao) validar() error {

	if publicacao.Titulo == "" {
		return errors.New("o título é obrigatório e não pode estar em branco")
	}
	if publicacao.Conteudo == "" {
		return errors.New("o conteúdo é obrigatório e não pode estar em branco")
	}
	return nil
}

// formatar é a função responsável por formatar os campos da publicação, removendo espaços em branco desnecessários.
func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
