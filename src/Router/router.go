package router

import (
	rotas "api/src/Router/Rotas"

	"github.com/gorilla/mux"
)

// Gerar retorna um router com todas rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
