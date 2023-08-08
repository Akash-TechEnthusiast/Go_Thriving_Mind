package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var person Person

		// Decode the JSON data from the request body into the Person struct
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		// You can now access the person's data and perform any desired actions
		fmt.Println("Received data:", person)

		fmt.Fprintf(w, "Received POST request successfully!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/postendpoint", handlePostRequest)
	http.ListenAndServe(":8080", nil)
}
