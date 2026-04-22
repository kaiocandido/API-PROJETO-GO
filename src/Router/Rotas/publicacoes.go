package rotas

import (
	controllers "api/src/Controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:        "/publicacoes",
		Metodo:     http.MethodPost,
		Func:       controllers.CriarPublicacao,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes",
		Metodo:     http.MethodGet,
		Func:       controllers.BuscarPublicacoes,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodGet,
		Func:       controllers.BuscarPublicacaoPorId,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodDelete,
		Func:       controllers.DeletarPublicacao,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodPut,
		Func:       controllers.AtualizarPublicacao,
		RequerAuth: true,
	},
	{
		URI:        "/usuarios/{usuarioId}/publicacoes",
		Metodo:     http.MethodGet,
		Func:       controllers.BuscarTodasPublicacoesPorUsuario,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}/curtir",
		Metodo:     http.MethodPost,
		Func:       controllers.CurtirPublicacao,
		RequerAuth: true,
	},
	{
		URI:        "/publicacoes/{publicacaoId}/deslike",
		Metodo:     http.MethodPost,
		Func:       controllers.DeslikePublicacao,
		RequerAuth: true,
	},
}
