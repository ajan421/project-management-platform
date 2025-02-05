package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/projects", GetProjects).Methods("GET")
	r.HandleFunc("/api/projects", CreateProject).Methods("POST")
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Fetching projects..."))
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a new project..."))
}
