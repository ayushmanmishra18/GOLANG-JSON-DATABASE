package main

import (
	"fmt"
	"net/http"
)

var db *DB // GLOBAL DB

func main() {
	var err error

	db, err = New("./data", nil) // store files in /data folder
	if err != nil {
		fmt.Println("DB error:", err)
		return
	}

	// Routes
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", userHandler)

	fmt.Println("🚀 Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}