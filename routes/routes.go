package routes

import (
	"backend-golang/controllers"
	"backend-golang/database"
	"backend-golang/repositories"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

type Response struct {
	Message string `json:"message"`
}

type ApiStatus struct {
	AppStatus    string `json:"app_status"`
	DBConnection string `json:"db_connection"`
}

// ConfiguraRotas configura as rotas e recebe o client do MongoDB
func ConfiguraRotas(client *mongo.Client) *mux.Router {

	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")

	repo, err := repositories.NovoPerfilRepositorio(dbURL, dbName)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Inicializar o controlador de perfil
	perfilController := controllers.NovoPerfilController(repo)

	router := mux.NewRouter()

	// Definir rotas
	router.HandleFunc("/perfis", perfilController.CriarPerfil).Methods("POST")
	router.HandleFunc("/perfis", perfilController.ListarTodosPerfis).Methods("GET")
	router.HandleFunc("/perfis/{id}", perfilController.ListarPerfilPorID).Methods("GET")
	router.HandleFunc("/perfis/{id}", perfilController.EditarPerfil).Methods("PUT")
	router.HandleFunc("/perfis/{id}", perfilController.DeletarPerfil).Methods("DELETE")

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		dbStatus := "503"
		if database.CheckConnection(client) {
			dbStatus = "200"
		}

		response := ApiStatus{
			AppStatus:    "200",
			DBConnection: dbStatus,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	return router
}
