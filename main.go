package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)



func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	// Your logic to handle the GET request goes here

	// In this example, we'll simply write "Hello, GET request!" as the response
	w.Write([]byte("Hello, GET request!"))
}



func handlePostWithPathParam(w http.ResponseWriter, r *http.Request) {
	// Get the path parameters from the request
	vars := mux.Vars(r)
	id := vars["id"]

	// Respond with the extracted "id" value
	fmt.Fprintf(w, "Received POST request with path parameter: %s", id)
}
func functionA() {
    fmt.Println("This is functionA from main.go")
}




func main() {



    // Defer closing the database connection

    //utils.SayHello()
    r := mux.NewRouter()
	r.HandleFunc("/post/{id}", handlePostWithPathParam).Methods("POST")
	// Start the HTTP server on port 8080 with the Gorilla Mux router
	http.Handle("/", r)
	fmt.Println("Server listening on port 8080")
	


   // http.HandleFunc("/", handleGetRequest)

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)


    fmt.Println("Hello, Go!")

    var student1 string = "John" //type is string
    var student2 = "Jane" //type is inferred
    x := 2 //type is inferred
  
    fmt.Println(student1)
    fmt.Println(student2)
    fmt.Println(x)

    var a string
    var b int
    var c bool
  
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

    var student3 string
    student3 = "John"
    fmt.Println(student3)

    var num int = 10


    for i := 1; i <= 2; i++ {
        fmt.Println(num)
        test(num)
        num -= 1
        fmt.Println(num)
    }
   


}

func test(num int) int {
    //var num int = 10
  
    num -= 1 // Decrement num by 1
    fmt.Println(num)
    return num // Output: 9
}