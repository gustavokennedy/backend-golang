package routes

import (
	"encoding/json"
	"net/http"

	"backend-golang/database"

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

	router := mux.NewRouter()

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
