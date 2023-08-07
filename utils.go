package main

import (
    "database/sql"
    "fmt"
   _ "github.com/go-sql-driver/mysql"
   "time"
   "strconv"
)

func main() {
    // Open the MySQL database connection
	//sql.Register("mysql", &mysql.MySQLDriver{})

    resultChan1 := make(chan string)
  //  resultChan2 := make(chan string)

     db, err := sql.Open("mysql", "root:KGM@123$@tcp(localhost:3306)/loadtest_kagamierp")

   // Check for errors
    if err != nil {
        fmt.Println("Error connecting to the database:", err)
        return
                 }else {
		fmt.Println("data base connected ")
	}
        startTime := time.Now()

        query := "SELECT id,status,empno,orgname FROM loadtestentity_39 WHERE status='Active'"
      //  quer1 := "SELECT id,status,empno,orgname FROM loadtestentity_39 WHERE status='Active1'"

     //   executeQuery(db,query,900000)

       go executeQuery(db,query,resultChan1,900000)


        data1 := <-resultChan1
       //	data2 := <-resultChan2

	     fmt.Println(data1)
      // fmt.Println(data2)
 
        endTime := time.Now()

        // Calculate the time difference
        elapsedTime := endTime.Sub(startTime)
        fmt.Printf("Time taken: %v\n", elapsedTime)
 

    // Print the time difference



    // Check for any errors during iteration
 
    // Defer closing the database connection
   // defer rows.Close()
    defer db.Close()
   



    // Now you can use the 'db' variable to perform database operations.
    // For example, you can execute queries, insert data, fetch data, et

}

func executeQuery(db *sql.DB, query string,resultChan chan<- string,b int ){

   // fmt.Printf(statusvalue)

    //query := "SELECT id,status,empno,orgname FROM loadtestentity_39 WHERE status='Active'"
    //fmt.Printf(statusvalue,"<->",query)
    rows, err := db.Query(query)
    if err != nil {
        panic(err.Error())
    }
    
    a:=b;
    // Iterate over the result set
    for rows.Next() {
        var column1Value string
        var column2Value string
        var column3Value string
        var column4Value string

   

        err := rows.Scan(&column1Value, &column2Value, &column3Value,&column4Value)
        if err != nil {
            panic(err.Error())
        }

        updateQuery := "UPDATE loadtestentity_39 SET status = 'kagoexecuted' WHERE id ="+column1Value+""
        result, err := db.Exec(updateQuery)
        if err != nil {
            panic(err.Error())
        }
    
        // Get the number of rows affected by the update
        numRowsAffected, err := result.RowsAffected()
        if err != nil {
            panic(err.Error())
        }
        a += 1 
        kk:= insertData(db,column2Value, column3Value,column4Value,a)
       
    
        fmt.Printf("Number of rows affected: %d %d\n", numRowsAffected,kk)
     
        //fmt.Println(kk)
    
       // fmt.Println(column1Value, column2Value, column3Value, a )
    }

    resultChan <- fmt.Sprintf("Data from %s", query)
   
}


func functionB(x, y int) int {
    return x + y
}

// functionA calls functionB and prints the result
func functionA() {
    result := functionB(10, 5)
    fmt.Println("Result of functionB:", result)
   // main.functionA() 
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