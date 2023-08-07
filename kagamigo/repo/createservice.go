package repo

import (
    "fmt"
)



func init() {
	fmt.Println("create service package initialized")
}
// functionA calls functionB and prints the result
func SomeFunction() string{
   // result := functionB(10, 5)
    return "create service called from main package"
   // main.functionA(uj) 
}




func ParentFunction(){
    // result := functionB(10, 5)
     //return "create service called from main package"

     fmt.Println("executing parent function")
    
    
 }
