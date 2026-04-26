package main

import (
	"encoding/json"
	"net/http"
)

// GET all users + POST user
func usersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		records, err := db.Readall("users")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(records)

	case "POST":
		var user User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid JSON", 400)
			return
		}

		if user.Name == "" {
			http.Error(w, "Name required", 400)
			return
		}

		if err := db.Write("users", user.Name, user); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte("User created"))

	default:
		http.Error(w, "Method not allowed", 405)
	}
}

// GET one / PUT / DELETE
func userHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[len("/users/"):]

	switch r.Method {

	case "GET":
		data, err := db.Read("users", id)
		if err != nil {
			http.Error(w, "User not found", 404)
			return
		}
		w.Write([]byte(data))

	case "PUT":
		var user User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid JSON", 400)
			return
		}

		if err := db.Update("users", id, user); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte("User updated"))

	case "DELETE":
		if err := db.Delete("users", id); err != nil {
			http.Error(w, "User not found", 404)
			return
		}

		w.Write([]byte("User deleted"))

	default:
		http.Error(w, "Method not allowed", 405)
	}
}