package main

import "fmt"

func quickSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    pivot := arr[0]
    var less, equal, greater []int

    for _, value := range arr {
        switch {
        case value < pivot:
            less = append(less, value)
        case value == pivot:
            equal = append(equal, value)
        case value > pivot:
            greater = append(greater, value)
        }
    }

    // Recursively sort the partitions
    less = quickSort(less)
    greater = quickSort(greater)

    // Concatenate the partitions back together
    return append(append(less, equal...), greater...)
}
func callquicksort(){

    fmt.Println("Sorted array:")
    arr := []int{64, 34, 25, 12, 22, 11, 90,89,67,56,78,45,23,90,54,22,55,66,33,99}
    fmt.Println("Unsorted array:", arr)

	 sortedArr1 :=quickSort(arr);
	//go sortedArr2 :=quickSort(arr);
   // sortedArr := quickSort(arr)
    fmt.Println("Sorted array:",sortedArr1)
}

func main() {

    go callquicksort()
    go callquicksort()

  
 
}