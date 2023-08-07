package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {

	jsonData := `{"name": "John Doe", "age": 30, "address": "123 Main St"}`
	var person Person

	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Address:", person.Address)
}