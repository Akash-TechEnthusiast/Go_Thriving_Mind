package service

import (
   
    "fmt"
	"kagamigo/service/utility"

   
  
)


func ServiceFunction() string{
    // result := functionB(10, 5)
   

     fmt.Println("executing service call")

	return  utility.Utilityfunction();
    // fmt.Println(controller.ControllerFunction())
    
    
 }
 