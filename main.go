package main

import (
	"fmt"
	"os"
	"encoding/json"
 ) 

 type Address struct {
	City string
	State string
	Country string
	PinCode string
 }
 type user struct{
	Name string
	Age json.Number
	Contact string
	Company string
	Address string
 }

func main() {
	println("Hello, World!")
}