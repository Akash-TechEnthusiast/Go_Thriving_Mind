package controller

import (
    "database/sql"
    "fmt"
   _ "github.com/go-sql-driver/mysql"
   "time"
   "kagamigo/service"
   "gorm.io/gorm"
   "gorm.io/driver/mysql"
   "kagamigo/model"

)

func main() {
    // Open the MySQL database connection
	//sql.Register("mysql", &mysql.MySQLDriver{})



     db, err := sql.Open("mysql", "root:KGM@123$@tcp(localhost:3306)/goschema")

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

       executeQuery(db,query,900000)


   
       //	data2 := <-resultChan2

	    // fmt.Println(data1)
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

func executeQuery(db *sql.DB, query string,b int ){

   // fmt.Printf(statusvalue)

    //query := "SELECT id,status,empno,orgname FROM loadtestentity_39 WHERE status='Active'"
    //fmt.Printf(statusvalue,"<->",query)
    rows, err := db.Query(query)
    if err != nil {
        panic(err.Error())
    }
    
   // a:=b;
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

        
    
       // fmt.Printf("Number of rows affected: %d %d\n", numRowsAffected,kk)
     
        //fmt.Println(kk)
    
       // fmt.Println(column1Value, column2Value, column3Value, a )
    }

  
   
}





func ControllerFunction() string{
    // result := functionB(10, 5)
     //return "create service called from main package"

     fmt.Println("executing controller call")
     return service.ServiceFunction()
    
    
 }


 func GenerateTables() string{

    fmt.Println("checkig the generate table")

    dsn := "root:KGM@123$@tcp(localhost:3306)/goschema?charset=utf8mb4&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    // AutoMigrate will create the "users" table in the MySQL database.
    db.AutoMigrate(&model.User{})
   // db.AutoMigrate(&model.Person{})

    db.AutoMigrate(&model.Purchaseorder{}, &model.Itemlines{})

    db.AutoMigrate(&model.Author{})
    db.AutoMigrate(&model.Post{})
    db.AutoMigrate(&model.PostUser{})
    

    return "db executing controller call"

 }
