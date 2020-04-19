package controllers

import "net/http"

// RegisterControllers : Register a new user controller if passed /users
func RegisterControllers() {

	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
