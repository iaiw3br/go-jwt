package main

import (
	"encoding/json"
	"fmt"
	"go-jwt/pkg/models"
	"log"
	"net/http"
)

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.SignUp
	err := decoder.Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	existUser := app.User.ExistUser(user.Name)

	if existUser {
		fmt.Fprintf(w, "Пользователь уже существует")
		return
	}

	hash, err := hashPassword(user.Password)

	if err != nil {
		log.Fatal(err)
	}

	user.Password = hash

	id, err := app.User.Create(user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, id)
}
