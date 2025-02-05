package main

import (
	"fmt"
	"log"
	"net/http"
	"project-management-platform/backend/config"

	"github.com/gorilla/mux" // Ensure this path is correct
)

func main() {
	// Connect to MongoDB
	config.ConnectDB()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Project Management Platform API!")
	}).Methods("GET")

	fmt.Println("🚀 Server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
