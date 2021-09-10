package main

import (
	mux2 "github.com/gorilla/mux"
)

func (app *application) routes() *mux2.Router {
	mux := mux2.NewRouter()
	mux.HandleFunc("/signup", app.signup).Methods("POST")
	mux.HandleFunc("/login", app.login).Methods("POST")

	return mux
}
