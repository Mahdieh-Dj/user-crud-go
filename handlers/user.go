package handlers

import (
	"encoding/json"
	"net/http"
	"user-api-go/models"
	"user-api-go/storage"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := storage.GetAllUsers()

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, found := storage.GetUser(id)

	if !found {
		http.Error(w, "User Not found", http.StatusNotFound)

		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "User Not found", http.StatusBadRequest)

		return
	}

	u.ID = uuid.NewString()
	storage.AddUser(u)

	json.NewEncoder(w).Encode(u)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	storage.DeleteUser(id)

	w.WriteHeader(http.StatusNoContent)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "User Not found", http.StatusBadRequest)

		return
	}

	storage.UpdateUser(id, u)

	json.NewEncoder(w).Encode(u)
}
