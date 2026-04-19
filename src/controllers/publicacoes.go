package controllers

import (
	autenticacao "api/src/Autenticacao"
	model "api/src/Model"
	"api/src/answers"
	"api/src/banco"
	"api/src/repository"
	"encoding/json"
	"io"
	"net/http"
)

// CriarPublicacao é a função responsável por criar uma nova publicação.
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := autenticacao.ExtrairUsuarioID(r)

	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	corpoReq, err := io.ReadAll(r.Body)

	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	var publicacao model.Publicacao

	if err = json.Unmarshal(corpoReq, &publicacao); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	publicacao.AutorID = usuarioID

	db, err := banco.Conectar()

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repository.NovoRepositorioPublicacoes(db)

	publicacaoID, err := repo.Criar(publicacao)

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	publicacao.ID = publicacaoID

	answers.JSON(w, http.StatusCreated, publicacao)

}

// BuscarPublicacoes é a função responsável por buscar todas as publicações.
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

// BuscarPublicacaoPorId é a função responsável por buscar uma publicação específica pelo ID.
func BuscarPublicacaoPorId(w http.ResponseWriter, r *http.Request) {
}

// DeletarPublicacao é a função responsável por deletar uma publicação específica pelo ID.
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
}

// AtualizarPublicacao é a função responsável por atualizar uma publicação específica pelo ID.
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
}
