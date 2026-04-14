package rotas

import (
	controllers "api/src/Controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:        "/usuarios",
		Metodo:     http.MethodPost,
		Func:       controllers.CriarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios",
		Metodo:     http.MethodGet,
		Func:       controllers.BuscarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios/{id}",
		Metodo:     http.MethodGet,
		Func:       controllers.BuscarUsuarioPorId,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios/{id}",
		Metodo:     http.MethodPut,
		Func:       controllers.AtualizarUsuario,
		RequerAuth: false,
	},
	{
		URI:        "/usuarios/{id}",
		Metodo:     http.MethodDelete,
		Func:       controllers.DeletarUsuario,
		RequerAuth: false,
	},
}
