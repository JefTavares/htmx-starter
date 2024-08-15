package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jeftavares/htmx-starter/views"
)

func loadUsers() ([]string, error) {
	res, err := http.Get(`https://fakerapi.it/api/v1/users?_quantity=3`)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var user struct {
		Data []struct {
			Firstname string `json:"firstname"`
		} `json:"data"`
	}
	json.NewDecoder(res.Body).Decode(&user)

	response := []string{}
	for _, u := range user.Data {
		response = append(response, u.Firstname)
	}

	return response, nil
}

func UsersHandler(w http.ResponseWriter, r *http.Request) error {
	users, err := loadUsers()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	time.Sleep(150 * time.Millisecond)
	return render(w, r, views.Users(users))
}

func UsersListHandler(w http.ResponseWriter, r *http.Request) error {
	users, err := loadUsers()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	time.Sleep(150 * time.Millisecond)
	return render(w, r, views.UsersList(users))
}
