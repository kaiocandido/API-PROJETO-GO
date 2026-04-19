package model

// senha representa a estrutura para atualizar a senha do usuario
type Senha struct {
	NovaSenha  string `json:"novaSenha"`
	SenhaAtual string `json:"senhaAtual"`
}
