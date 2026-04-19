package controllers

import (
	autenticacao "api/src/Autenticacao"
	model "api/src/Model"
	"api/src/answers"
	"api/src/banco"
	"api/src/repository"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuario envia um usuario para o repository
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)

	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario model.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := usuario.Preparar("cadastro"); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioUsuarios(db)
	usuarioId, err := repositorio.Criar(usuario)

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	usuario.ID = usuarioId

	answers.JSON(w, http.StatusCreated, usuario)

}

// BuscarUsuario procura um usuario
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := banco.Conectar()

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repository.NovoRepositorioUsuarios(db)

	usuarios, err := repo.Buscar(nomeOuNick)

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, usuarios)
}

// BuscarUsuarioPorId procura um usuario pelo ID unico
func BuscarUsuarioPorId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()
	repo := repository.NovoRepositorioUsuarios(db)
	usuario, err := repo.BuscarPorId(usuarioId)

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, usuario)
}

// AtualizarUsuario faz a atualização atraves do ID do usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametro["usuarioId"], 10, 64)

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	usuarioIdToken, err := autenticacao.ExtrairUsuarioID(r)

	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIdToken {
		answers.Erro(w, http.StatusForbidden, errors.New("não é permitido atualizar um usuario diferente do seu"))
		return
	}

	corpoReq, err := ioutil.ReadAll(r.Body)

	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario model.Usuario

	if err = json.Unmarshal(corpoReq, &usuario); err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = usuario.Preparar("edicao"); err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := banco.Conectar()

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioUsuarios(db)

	if err = repositorio.Atualizar(usuarioID, usuario); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)

}

// DeletarUsuario exclui um usuario
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioId, err := strconv.ParseUint(parametros["usuarioID"], 10, 64)

	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	usuarioIdToken, err := autenticacao.ExtrairUsuarioID(r)

	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioId != usuarioIdToken {
		answers.Erro(w, http.StatusForbidden, errors.New("não é permitido deletar um usuario diferente do seu"))
		return
	}

	db, err := banco.Conectar()

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioUsuarios(db)

	if err = repositorio.Deletar(usuarioId); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)

}

// SeguirUsuario permite que um usuario siga outro usuario
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorId, err := autenticacao.ExtrairUsuarioID(r)

	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)

	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if seguidorId == usuarioId {
		answers.Erro(w, http.StatusForbidden, errors.New("não é permitido seguir você mesmo"))
		return
	}

	db, err := banco.Conectar()

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioUsuarios(db)

	if err = repositorio.Seguir(usuarioId, seguidorId); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)

}

// PararDeSeguirUsuario permite que um usuario deixe de seguir outro usuario
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorId, err := autenticacao.ExtrairUsuarioID(r)

	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)

	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if usuarioId == seguidorId {
		answers.Erro(w, http.StatusForbidden, errors.New("não é permitido deixar de seguir você mesmo"))
		return
	}

	db, err := banco.Conectar()

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioUsuarios(db)

	if err = repositorio.PararDeSeguirUsuario(usuarioId, seguidorId); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)

}

// BuscarSeguidores traz todos os seguidores de um usuario
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := banco.Conectar()

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioUsuarios(db)

	seguidores, err := repositorio.BuscarSeguidores(usuarioId)

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, seguidores)
}

// BuscarSeguindo traz os usuarios que um usuario segue
func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := banco.Conectar()

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorio := repository.NovoRepositorioUsuarios(db)

	usuarios, err := repositorio.BuscarSeguindo(usuarioId)

	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, usuarios)
}

// AtualizarSenha permite que um usuario atualize sua senha
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIdToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if usuarioId != usuarioIdToken {
		answers.Erro(w, http.StatusForbidden, errors.New("não é permitido atualizar a senha de outro usuário"))
		return
	}

	corpoReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusServiceUnavailable, err)
		return
	}

	var senha model.Senha
	if err = json.Unmarshal(corpoReq, &senha); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioUsuarios(db)

	// Busca a senha salva no banco de dados
	senhaSalvaNoBanco, err := repositorio.BuscarSenha(usuarioId)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	// Verifica se a senha atual fornecida corresponde à senha salva no banco
	if err = security.VerificarSenha(senha.Atual, senhaSalvaNoBanco); err != nil {
		answers.Erro(w, http.StatusUnauthorized, errors.New("senha incorreta"))
		return
	}

	// Gera o hash da nova senha
	senhaHash, err := security.Hash(senha.Nova)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	// Atualiza a senha no banco de dados
	if err = repositorio.AlterarSenha(usuarioId, string(senhaHash)); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	// Retorna uma resposta de sucesso
	answers.JSON(w, http.StatusOK, nil)
}
