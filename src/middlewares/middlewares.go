package middlewares

import (
	autenticacao "api/src/Autenticacao"
	"api/src/answers"
	"log"
	"net/http"
)

// Logger é um middleware que registra as informações da requisição, como método, URI e host.
func Logger(proximaFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFunc(w, r)
	}
}

// Autenticar é um middleware que verifica se o usuário está autenticado antes de permitir o acesso a uma rota protegida.
func Autenticar(proximaFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := autenticacao.ValidarToken(r); err != nil {
			answers.Erro(w, http.StatusUnauthorized, err)
			return
		}
		proximaFunc(w, r)
	}
}
