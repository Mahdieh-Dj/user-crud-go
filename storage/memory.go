package storage

import (
	"sync"
	"user-api-go/models"
)

var users = make(map[string]models.User)

var mu sync.Mutex

func AddUser(u models.User) {
	mu.Lock()
	defer mu.Unlock()

	users[u.ID] = u
}

func GetUser(id string) (models.User, bool) {
	mu.Lock()
	defer mu.Unlock()

	user, found := users[id]

	return user, found
}

func GetAllUsers() []models.User {
	mu.Lock()
	defer mu.Unlock()

	result := make([]models.User, 0, len(users))

	for _, u := range users {
		result = append(result, u)
	}

	return result
}

func DeleteUser(id string) {
	mu.Lock()
	defer mu.Unlock()

	delete(users, id)
}

func UpdateUser(id string, u models.User) models.User {
	users[id] = u

	return u
}
