package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var dbURL string
var dbName string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Falha ao carregar o arquivo .env.")
	}

	ambiente := os.Getenv("AMBIENTE")
	if ambiente == "dev" {
		log.Println("Configurando banco de dados para Desenvolvimento.")
		dbURL = "mongodb://mongo:27017"
		dbName = "backend-golang"
	} else if ambiente == "prod" {
		log.Println("Configurando banco de dados para Produção.")
		dbURL = "mongodb://mongo:@"
		dbName = "backend-golang"
	} else {
		log.Println("Valor de ambiente inválido. Usando valores padrão do Docker.")
		dbURL = "mongodb://mongo:27017"
		dbName = "backend-golang"
	}
}

// GetDBURL retorna a URL do banco de dados
func GetDBURL() string {
	return dbURL
}

// GetDBName retorna o nome do banco de dados
func GetDBName() string {
	return dbName
}