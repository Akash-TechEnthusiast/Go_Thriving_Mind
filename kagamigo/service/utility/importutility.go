package utility

import (
   
    "fmt"
    "kagamigo/repo"
)


func Utilityfunction() string{
    // result := functionB(10, 5)
    fmt.Println("executing utility call")
    repo.ParentFunction()

    return "returned from utility "
 
    
 
}
 