package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address string
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error creating database:", err)
		return
	}

	employee := []User{
		{"Ayushman", json.Number("30"), "ayushman@example.com", "ABC Corp", "Delhi"},
		{"Rohit", json.Number("25"), "rohit@example.com", "XYZ Corp", "Gurgaon"},
	}

	// Write
	for _, val := range employee {
		if err := db.Write("users", val.Name, val); err != nil {
			fmt.Println("Write error:", err)
		}
	}

	// Update
	updatedUser := User{
		"Ayushman", json.Number("35"),
		"ayushman_new@example.com", "ABC Corp", "Delhi",
	}

	if err := db.Update("users", "Ayushman", updatedUser); err != nil {
		fmt.Println("Update error:", err)
	}

	// Delete
	if err := db.Delete("users", "Rohit"); err != nil {
		fmt.Println("Delete error:", err)
	} else {
		fmt.Println("Rohit deleted")
	}

	// Concurrent writes
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			user := User{
				Name:    fmt.Sprintf("User%d", i),
				Age:     json.Number("20"),
				Contact: "test@example.com",
				Company: "TestCorp",
				Address: "City",
			}

			if err := db.Write("users", user.Name, user); err != nil {
				fmt.Println("Concurrent write error:", err)
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("All concurrent writes done")

	// Final read
	records, _ := db.Readall("users")
	fmt.Println("Final Records:", records)
}