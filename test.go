package main

import "fmt"

func main() {

    content:="test text,"
    createtablesql := content[:len(content)-1]
    fmt.Println(createtablesql)

  
 
}