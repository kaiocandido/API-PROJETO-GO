package rotas

import (
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

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Func).Methods(rota.Metodo)
	}

	return r

}
