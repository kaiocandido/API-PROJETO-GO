package controllers

import (
	model "api/src/Model"
	"api/src/answers"
	"api/src/banco"
	"api/src/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	w.Write([]byte("Buscando usuario"))
}

// BuscarUsuarioPorId procura um usuario pelo ID unico
func BuscarUsuarioPorId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuario por ID"))
}

// AtualizarUsuario faz a atualização atraves do ID do usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario"))
}

// DeletarUsuario exclui um usuario
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario"))
}
