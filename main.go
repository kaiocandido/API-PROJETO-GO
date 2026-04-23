package main

import (
	router "api/src/Router"
	"api/src/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	config.Carregar()
	fmt.Println("Rodando API")

	r := router.Gerar()

	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:5173"})

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", config.Porta),
		handlers.CORS(headersOk, methodsOk, originsOk)(r),
	))
}
