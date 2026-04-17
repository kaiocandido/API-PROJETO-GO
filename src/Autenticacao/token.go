package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/dgrijalva/jwt-go"
	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken Gera um token com as permissões
func CriarToken(usuarioId uint64) (string, error) {
	permisoes := jwt.MapClaims{}

	permisoes["authorized"] = true
	permisoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permisoes["usuarioId"] = usuarioId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permisoes)
	return token.SignedString([]byte(config.Key))
}

// ValidarToken Verifica se o token é válido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)

	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

// extrairToken Extrai o token da requisição
func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// retornarChaveDeVerificacao Retorna a chave de verificação do token
func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, http.ErrAbortHandler
	}

	return config.Key, nil
}

// ExtrairUsuarioID Extrai o ID do usuário do token
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)

	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioId, err := strconv.ParseUint(fmt.Sprintf("%0.f", permissoes["usuarioId"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return usuarioId, nil
	}

	return 0, errors.New("token inválido")

}
