package main

//import "fmt"

func getSequence(i int) func() int {
    return func() int {
        i++
        return i
    }
}

func main(){
    nextNumber := getSequence(10)
    nextNumber1 := getSequence(100)
    println(nextNumber())
    println(nextNumber())
    println(nextNumber())
    println(nextNumber1())
    println(nextNumber1())
}

