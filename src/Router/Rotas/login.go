package rotas

import (
	controllers "api/src/Controllers"
	"net/http"
)

var rotaLogin = Rota{
	URI:        "/login",
	Metodo:     http.MethodPost,
	Func:       controllers.Login,
	RequerAuth: false,
}
