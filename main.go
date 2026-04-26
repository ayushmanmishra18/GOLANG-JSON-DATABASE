package main

import (
	"encoding/json"
	"fmt"
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
		{
			Name:    "Ayushman",
			Age:     json.Number("30"),
			Contact: "ayushman@example.com",
			Company: "ABC Corp",
			Address: "Delhi",
		},
		{
			Name:    "Rohit",
			Age:     json.Number("25"),
			Contact: "rohit@example.com",
			Company: "XYZ Corp",
			Address: "Gurgaon",
		},
	}

	for _, val := range employee {
		if err := db.Write("users", val.Name, val); err != nil {
			fmt.Println("Write error:", err)
		}
	}

	records, err := db.Readall("users")
	if err != nil {
		fmt.Println("Error reading from database:", err)
		return
	}

	allUsers := []User{}
	for _, f := range records {
		employeefound := User{}
		if err := json.Unmarshal([]byte(f), &employeefound); err != nil {
			fmt.Println("Error unmarshalling:", err)
			continue
		}
		allUsers = append(allUsers, employeefound)
	}

	fmt.Println("Parsed Users:", allUsers)

// Read single record
	data, err := db.Read("users", "Ayushman")
if err != nil {
	fmt.Println("Read error:", err)
} else {
	fmt.Println("Single Record:", data)
}


//update record
updatedUser := User{
	Name:    "Ayushman",
	Age:     json.Number("35"), // updated age
	Contact: "ayushman_new@example.com",
	Company: "ABC Corp",
	Address: "Delhi",
}

db.Update("users", "Ayushman", updatedUser)

//delete record
err = db.Delete("users", "Rohit")
if err != nil {
	fmt.Println("Delete error:", err)
} else {
	fmt.Println("Rohit deleted")
}
}