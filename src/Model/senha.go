package model

// senha representa a estrutura para atualizar a senha do usuario
type Senha struct {
	Nova  string `json:"novaSenha"`
	Atual string `json:"senhaAtual"`
}
