package controllers

import (
	autenticacao "api/src/Autenticacao"
	model "api/src/Model"
	"api/src/answers"
	"api/src/banco"
	"api/src/repository"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login é a função responsável por lidar com as requisições de login
func Login(w http.ResponseWriter, r *http.Request) {
	corpoReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario model.Usuario

	if err = json.Unmarshal(corpoReq, &usuario); err != nil {
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

	usuarioSalvoNoBanco, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerificarSenha(usuario.Senha, usuarioSalvoNoBanco.Senha); err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
