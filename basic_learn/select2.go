package main


import "time"
import "fmt"

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(time.Second*1)
        c1 <- "one"
    }()

    go func() {
        time.Sleep(time.Second*2)
        c2 <- "two"
    }()

    for i:=0;i<2;i++ {
        select {
            case msg := <- c1:
                fmt.Printf("c1 received: %s\n", msg)
            case msg := <- c2:
                fmt.Printf("c2 received: %s\n", msg)
        }
    }

}


