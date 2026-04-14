package model

import "time"

// Usuario representa um usuario, e seus dados
type Usuario struct {
	ID     uint64    `json:"id,omitempty"`
	Nome   string    `json:"nome,omitempty"`
	Nick   string    `json:"nick,omitempty"`
	Email  string    `json:"email,omitempty"`
	Senha  string    `json:"senha,omitempty"`
	Criado time.Time `json:"criado,omitempty"`
}
