package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"backend-golang/routes"
	// "backend-golang/seeders"

	"github.com/getsentry/sentry-go"
	"github.com/gorilla/handlers"
)

func main() {
	// Verifica se o diretório "logs" existe, caso contrário, cria-o
	err := os.Mkdir("logs", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	// Cria o arquivo de log
	logFile, err := os.Create("logs/server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Configura o logger para usar o arquivo de log
	log.SetOutput(logFile)
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multiWriter)

  // Sentry
	err = sentry.Init(sentry.ClientOptions{
		Dsn:              "",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second) 
	//sentry.CaptureMessage(Funcionando!")
  
	// Cria um roteador principal com Mux
	router := routes.ConfiguraRotas()

	// Configuração personalizada do CORS
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)



	// Define a porta do servidor
	port := ":8080"

	// Inicia o servidor com CORS
	// Precisa incluir router no cors cors(router)
	log.Printf("Servidor rodando em http://localhost%s\n", port)
	if err := http.ListenAndServe(port, cors(router)); err != nil {
		log.Printf("Erro ao iniciar servidor: %v\n", err)
	}
}
