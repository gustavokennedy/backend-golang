package routes

import (
	"encoding/json"
	// "fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"backend-golang/config"
	"backend-golang/database"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}
type ApiStatus struct {
	Status string `json:"status_app"`
	DbCheck string `json:"status_db"`
}

func ConfiguraRotas() *mux.Router {

	dbURL := config.GetDBURL()
	dbName := config.GetDBName()

	// Tenta se conectar ao banco de dados
	db, err := database.Connect(dbURL, dbName)
	dbStatus := "500"
	if err != nil || db == nil {
		log.Println("❌ Erro ao conectar ao banco de dados:", err)
	} else {
		dbStatus = "200"
		log.Println("✅ Conexão com o banco de dados estabelecida com sucesso!")
	}

	// Repositories
	// Controladores

	router := mux.NewRouter()

	///  Rotas Liberadas

	// Rota para verificar o status da API e do banco de dados
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		response := ApiStatus{
			Status:  "200",
			DbCheck: dbStatus,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Rotas Protegidas

	///////// Perfil
	///////// Usuários


	return router
}