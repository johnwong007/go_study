package main

import "fmt"

func main(){
    var x interface{}
    //x = nil 
    switch i:=x.(type){
        case nil:
            fmt.Printf("x :%T", i)
        case float64:
            fmt.Printf("x is float64")
        case func(int) float64:
            fmt.Printf("x is func(int)")
        default:
            fmt.Printf("未知🌟")
    }
}

