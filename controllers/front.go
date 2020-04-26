package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers : Register a new user controller if passed /users
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
