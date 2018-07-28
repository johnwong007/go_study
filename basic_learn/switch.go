package main

import "fmt"

func main(){
    println("switch和c++中的不太一样，case语句自带break")
    var days = 1
    switch days {
        case 1:
            println("days is only 1")
        case 2:
            fmt.Println("days is 2")
        default:
            fmt.Println("other digits")
    }
}


