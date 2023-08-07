package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sharedResource int
	mu             sync.Mutex // Mutex to synchronize access to the shared resource
)

// worker is a function that increments the sharedResource without synchronization.
func worker(wg *sync.WaitGroup) {
	// Simulate some work before accessing the shared resource
	mu.Lock() 
	time.Sleep(100 * time.Millisecond)

	// Access and modify the shared resource without synchronization
	sharedResource++
	mu.Unlock()
	fmt.Printf("Shared resource updated: %d\n", sharedResource)
	
	wg.Done() // Signal that the work is done by decrementing the WaitGroup
}

func main() {
	const numWorkers = 5

	wg := sync.WaitGroup{} // WaitGroup to wait for all goroutines to finish
	wg.Add(numWorkers)

	// Start multiple workers
	for i := 0; i < numWorkers; i++ {
		go worker(&wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All goroutines finished.")
}
