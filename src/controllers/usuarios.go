package controllers

import "net/http"

// CriarUsuario insere um novo usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario"))
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
