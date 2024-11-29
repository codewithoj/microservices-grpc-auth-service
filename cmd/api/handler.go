package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (app *Config) CreateUser(w http.ResponseWriter, r *http.Request) {

	

	data := &User{
		Name:  "John Doe",
		Email: "john@email.com",
	}

	out, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(out)
	if err != nil {
		log.Println(err)
	}

}
