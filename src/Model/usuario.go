package model

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		return errors.New("O email inserido é invalido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatorio e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, err := security.Hash(usuario.Senha)
		if err != nil {
			return err
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}

// Preparar chama os metodos para validar o usuario recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}

	if err := usuario.formatar(etapa); err != nil {
		return err
	}

	return nil
}
