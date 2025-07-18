package main

import (
	"log"
	"net/http"
	"user-api-go/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
