package main

import (
	"fmt"
	"time"
)

// Function to make a call and return the result through a channel


func callAPI(url string, resultChan chan<- string) {
	// Simulate a call to the API by sleeping for some time
	
	fmt.Println("Hello, Go!")
	time.Sleep(time.Second * 2)
	// Return the result through the channel
	resultChan <- fmt.Sprintf("Data from %s", url)
}


func main() {
	// Create channels to receive results
	resultChan1 := make(chan string)
    resultChan2 := make(chan string)

	startTime := time.Now()
	go callAPI("https://api.example.com/endpoint1", resultChan1)
	//go callAPI("https://api.example.com/endpoint2", resultChan2)
  // Wait for both goroutines to finish and receive the results
	data1 := <-resultChan1
//	data2 := <-resultChan2

	fmt.Println(data1)
//	fmt.Println(data2)
	endTime := time.Now()

	// Calculate the time difference
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Time taken: %v\n", elapsedTime)




}

