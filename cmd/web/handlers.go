package main

import (
	"encoding/json"
	"fmt"
	"go-jwt/pkg/models"
	"log"
	"net/http"
)

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromBody(r)

	if err != nil {
		log.Fatal(err)
	}

	existUser := app.User.ExistUser(user.Name)

	if existUser {
		fmt.Fprintf(w, "Пользователь уже существует")
		return
	}

	id, err := app.User.Create(user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, id)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.UserAuth
	err := decoder.Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	foundUser, err := app.User.FindUserByName(user.Name)

	if err != nil {
		log.Fatal(err)
		return
	}

	isNotValidPassword := !checkPassword(user.Password, foundUser.Password)
	if isNotValidPassword {
		log.Fatal("Пароль указан не верно")
		return
	}

	fmt.Fprintf(w, "user is valid")
}

func getUserFromBody(r *http.Request) (models.UserAuth, error) {
	decoder := json.NewDecoder(r.Body)
	var user models.UserAuth
	err := decoder.Decode(&user)

	if err != nil {
		return user, err
	}

	hash, err := hashPassword(user.Password)

	if err != nil {
		log.Fatal(err)
	}

	user.Password = hash

	return user, nil
}
