package main

import (
	"fmt"
	"sync"
)

var (
	sharedResource int
	mu             sync.Mutex // Mutex to synchronize access to the shared resource
	value int
)

// worker is a function that increments the sharedResource in a synchronized manner.
func worker(wg *sync.WaitGroup, ch chan struct{}) {
	mu.Lock() // Lock the Mutex before accessing the shared resource
	sharedResource++
	value++
	fmt.Printf("Shared resource updated: %d,%d\n", sharedResource,value)
	mu.Unlock() // Unlock the Mutex

	wg.Done()   // Signal that the work is done by decrementing the WaitGroup
	ch <- struct{}{} // Send a message on the channel to indicate completion
}

func main() {
	const numWorkers = 5

	wg := sync.WaitGroup{} // WaitGroup to wait for all goroutines to finish
	wg.Add(numWorkers)

	ch := make(chan struct{}) // Channel to synchronize the goroutines

	// Start multiple workers
	for i := 0; i < numWorkers; i++ {
		go worker(&wg, ch)
	}

	// Wait for all workers to finish
	wg.Wait()

	// Receive messages from the channel to ensure all goroutines have completed
	for i := 0; i < numWorkers; i++ {
		<-ch
	}

	fmt.Println("All goroutines finished.")
}