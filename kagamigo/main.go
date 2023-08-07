package main

import (
   
    "fmt"

    // "kagamigo/repo"
    "kagamigo/controller"
  //  "kagamigo/model"
  
)

func main() {

    fmt.Println("calling from main function")
	fmt.Println(controller.ControllerFunction())
    fmt.Println("called")
    fmt.Println(controller.GenerateTables())
 //   model.create()
    //controller.generateTables()

}

 


