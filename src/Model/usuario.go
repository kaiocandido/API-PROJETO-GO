package model

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuario, e seus dados
type Usuario struct {
	ID     uint64    `json:"id,omitempty"`
	Nome   string    `json:"nome,omitempty"`
	Nick   string    `json:"nick,omitempty"`
	Email  string    `json:"email,omitempty"`
	Senha  string    `json:"senha,omitempty"`
	Criado time.Time `json:"criado,omitempty"`
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatorio e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatorio e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O email é obrigatorio e não pode estar em branco")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatorio e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}

// Preparar chama os metodos para validar o usuario recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}

	usuario.formatar()
	return nil
}
