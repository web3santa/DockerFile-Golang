package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users []User

func init() {
	users = []User{
		{1, "Devid"},
		{2, "Henry"},
		{3, "Alex"},
	}
}

func main() {
	uh := userHandler{}
	http.Handle("/users", uh)
	http.ListenAndServe(":8080", nil)
}

type userHandler struct{}

func (uh userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getUsers(w, r)
	default:
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
