package main

import (
    "database/sql"
    "fmt"
   // "os"

    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"
	//"strings"
)

const (
    host = "localhost"
    port = 5432
    user = "postgres"
    password = "KGM@123$"
    dbname = "postgres"
	schema = "next_gen_app"
)


func main() {
    // Get the MySQL connection string from the environment variable.
	db, err := sql.Open("mysql", "root:KGM@123$@tcp(localhost:3306)/ShilpaCare_kagamierp")
	if err != nil {
		//log.Fatal(err)
	}
	defer db.Close()

    // Get the name of the MySQL table to convert.
   // mysql_table_name1 := os.Getenv("accesstype_301")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    // Open a connection to the database.
     pg_db, err := sql.Open("postgres", psqlInfo)



    if err != nil {
        panic(err)
    }
    defer pg_db.Close()

	mysql_table_name:="accesstype_301"
    // Get the MySQL table schema.


	rows, err := db.Query("SHOW COLUMNS FROM `" + mysql_table_name + "`")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    // Iterate over the rows in the result set and print the column name and data type.
    for rows.Next() {
        var column_name string
        var data_type string
        err := rows.Scan(&column_name, &data_type)
        if err != nil {
           // panic(err)
        }

        fmt.Printf("Column name: %s Data type: %s\n", column_name, data_type)
    }

    
    fmt.Println("Successfully converted MySQL table to PostgreSQL table.")
}