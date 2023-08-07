package main

import (
	"database/sql"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
	

)

func main() {
    // Connect to the database
	db, err := sql.Open("mysql", "root:KGM@123$@tcp(localhost:3306)/loadtest_kagamierp")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Execute the SQL query

	// Step 2: Prepare and execute your SELECT query. id,status,empno,orgname
	query := "SELECT * FROM loadtestentity_39 WHERE status='Active'"
	rows, err := db.Query(query)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    // Get the number of rows in the result set
    numRows := 0
    for rows.Next() {
        numRows++
    }

	fmt.Println(numRows)
    // Calculate the number of batches
    batchSize := 100
    numBatches := numRows / batchSize
    if numRows % batchSize != 0 {
        numBatches++
    }


    // Create a slice to store the batches
    batches := make([][][]byte, 1)
	fmt.Println(batches)

    // Iterate over the result set and add the rows to the batches
    i := 0
    for rows.Next() {
        var row []byte
        err := rows.Scan(&row)
        if err != nil {
            panic(err)
        }
        batches[i%numBatches] = append(batches[i%numBatches], row)
		fmt.Println(batches[0])
        i++
    }

    // Print the batches
    for i := 0; i < numBatches; i++ {
        fmt.Println(batches[0])
		//fmt.Println(batches)
 
	}
}