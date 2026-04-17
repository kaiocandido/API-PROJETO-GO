package security

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma senha e retorna a senha criptografada
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara a senha recebida com a senha criptografada
func VerificarSenha(senha, senhaHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senha))
}
