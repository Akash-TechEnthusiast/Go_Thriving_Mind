
package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"strings"
	"strconv"

)
var (
	mu             sync.Mutex // Mutex to synchronize access to the shared resource
)
var sharedVariable int
const batchSize = 1000   // Desired batch size
const numWorkers = 4   // Number of worker goroutines

// processBatch processes a single batch of data.
func processBatch(db *sql.DB,workerID int, batch []string) {
	for _, item := range batch {
		// Process each item in the batch
		resultstring := strings.Split(item, "&")
		//fmt.Printf("Worker %d: %s\n", workerID, resultstring)
		updateQuery := "UPDATE loadtestentity_39 SET status = 'kagoexecuted' WHERE id ="+resultstring[0]+""
        result, err := db.Exec(updateQuery)
        if err != nil {
            panic(err.Error())
        }
    
        // Get the number of rows affected by the update
        numRowsAffected, err := result.RowsAffected()
        if err != nil {
            panic(err.Error())
        }
		
		mu.Lock() 
		sharedVariable++
        kk:= insertData(db,resultstring[1],resultstring[2],resultstring[3],sharedVariable)
		fmt.Printf("Number of rows affected:%d %d %d\n", workerID,numRowsAffected,kk)
		mu.Unlock()

       // fmt.Printf("Number of rows affected: %d %d\n", numRowsAffected,kk)
       
	}
}






func main() {
	// Step 1: Set up your database connection.
	sharedVariable=900000
	db, err := sql.Open("mysql", "root:KGM@123$@tcp(localhost:3306)/loadtest_kagamierp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Step 2: Prepare and execute your SELECT query. id,status,empno,orgname
	query := "SELECT id,status,empno,orgname FROM loadtestentity_39 WHERE status='Active'"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Step 3: Create a function to process a single batch of data.

	// Step 4: Divide the fetched data into smaller batches.
	// Step 3: Create a function to process a single batch of data.

	// Step 4: Divide the fetched data into smaller batches.


	//batches12 := divideRowsIntoBatches(rows, batchSize)
	startTime := time.Now()


	batch := make([]string, 0, batchSize)
	batches := make([][]string, 0)

	for rows.Next() {

		var column1Value string
        var column2Value string
        var column3Value string
        var column4Value string

		if err := rows.Scan(&column1Value, &column2Value, &column3Value,&column4Value); err != nil {
		//	if err := rows.Scan(&row); err != nil {
			log.Fatal(err)
		}
		str := column1Value+"&"+column2Value+"&"+ column3Value+"&"+column4Value
		batch = append(batch,str)
	//	fmt.Printf("batch data %s\n", batch)
	//	fmt.Printf("batch data %d\n", len(batch))
		if len(batch) == batchSize {
			batches = append(batches, batch)
			batch = make([]string, 0, batchSize)
		}
	}

	if len(batch) > 0 {
		batches = append(batches, batch)
	}

	// Step 5: Set up a channel to communicate between the main goroutine and the worker goroutines.
	workerChannel := make(chan struct {
		id    int
		batch []string
	})

	// Step 6: Launch worker goroutines that will process the batches concurrently.
	var wg sync.WaitGroup
	wg.Add(numWorkers) //numWorkers

	for i := 0; i < numWorkers; i++ { //numWorkers
		go func(workerID int) {
			defer wg.Done()
			for data := range workerChannel {
				// Step 7: Send the batches to the worker goroutines through the channel.
				processBatch(db,data.id, data.batch)
			}
		}(i)
	}

	// Step 8: Send the batches to the worker goroutines through the channel.
	for i, batch := range batches {
		workerChannel <- struct {
			id    int
			batch []string
		}{id: i, batch: batch}
	}
	close(workerChannel)

	// Step 9: Wait for all worker goroutines to finish processing.
	wg.Wait()

	 
	endTime := time.Now()

	// Calculate the time difference
	fmt.Printf("no of workers: %s\n", numWorkers)
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Time taken: %v\n", elapsedTime)
}


func insertData(db *sql.DB,status string ,empno string ,orgname string ,number int ) error {
    //db, err := createConnection()
  
    numStr := strconv.Itoa(number)
    insertQuery := "INSERT INTO loadtestentity_39 (id,status,empno,orgname,transactionentityid) VALUES (?, ?, ?,?,?)"

    stmt, err := db.Prepare(insertQuery)
	if err != nil {
        panic(err.Error())
	}
	defer stmt.Close()

	// Execute the INSERT statement with the provided data
	_, err = stmt.Exec(numStr,"kagoexecuted",empno,orgname,number)
	if err != nil {
        panic(err.Error())
	}
    return nil
}