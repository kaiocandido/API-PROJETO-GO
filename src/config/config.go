package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexãoBanco usado para conexão com mysql
	StringConexaoBanco = ""

	// Porta indica a porta onde a aplicação vai rodar
	Porta = 0

	// Senha do JWT
	Key []byte
)

// Carregar inicializa as variaveis de ambiente
func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NOME"),
	)

	Key = []byte(os.Getenv("KEY"))

}
