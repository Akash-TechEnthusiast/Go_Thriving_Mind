package model

//import "fmt"

// Define the embedded entity
type Address struct {
    Street  string
    City    string
    Country string
}

// Define the main entity with the embedded entity
type Person struct {
    Name    string
    Age     int
    Address []*Address
}
