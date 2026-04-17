package autenticacao

import (
	"api/src/config"
	"errors"
	"net/http"
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

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, http.ErrAbortHandler
	}

	return config.Key, nil
}
