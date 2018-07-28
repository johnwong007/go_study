package main

import "fmt"

func main() {
    var c1,c2,c3 chan int
    var i2 = 0
    select {
        case i1 := <-c1:
            fmt.Printf("received %d from c1", i1)
        case c2 <- i2:
            fmt.Printf("%d sent to c2", i2)
        case _,ok := (<-c3):
            if ok {
                fmt.Printf("c3 is ok")
            } else {
                fmt.Printf("c3 is not ok")
            }
        default:
            fmt.Printf("no communication")
    }

}

