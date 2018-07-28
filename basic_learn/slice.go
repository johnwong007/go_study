package main

import "fmt"

func main(){
    var slice1 []int = []int{0,0};
    printNumbers(slice1)
    var numbers = make([]int, 3, 5)
    printNumbers(numbers)
    var slice2 [10]int;
    fmt.Printf("len=%d cap=%d slice=%+v \n", len(slice2), cap(slice2), slice2)
    var slice3 = [3]int{1,2,3};
    fmt.Printf("len=%d cap=%d slice=%+v \n", len(slice3), cap(slice3), slice3)
}

func printNumbers(x []int){
    fmt.Printf("len=%d cap=%d slice=%+v \n", len(x), cap(x), x)

}


