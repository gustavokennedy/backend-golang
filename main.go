package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"backend-golang/routes"
	"backend-golang/seeders"

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
		Dsn:              "https://14e537b912b04f53848b3299d5eab886@o4504888820695040.ingest.sentry.io/4505365111570432",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second) 
	//sentry.CaptureMessage(Funcionando!")
  
	// Cria um roteador principal com Mux
	// router := routes.ConfiguraRotas()

	// Configuração personalizada do CORS
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Inicia servidor HTTP
	// log.Println("Servidor rodando na porta 8080...")
	// log.Fatal(http.ListenAndServe(":8080", cors(router)))

	// Define a porta do servidor e inicia
	port := ":8080"
	fmt.Printf("Servidor rodando em http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Erro ao iniciar servidor: %v\n", err)
	}
}
