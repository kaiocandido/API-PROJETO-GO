package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas rotas da API
type Rota struct {
	URI        string
	Metodo     string
	Func       func(http.ResponseWriter, *http.Request)
	RequerAuth bool
}

// Configurar coloca todas rotas
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasPublicacoes...)

	for _, rota := range rotas {
		if rota.RequerAuth {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Func))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Func)).Methods(rota.Metodo)
		}

	}

	return r

}
