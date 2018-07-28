package main

import "time"
import "fmt"

func main() {
    c1 := make(chan string)

    go func() {
        time.Sleep(time.Second*2)
        c1<-"c1 send"
    }()

    select {
        case msg := <-c1:
            fmt.Printf("Received msg:%s\n", msg)
        case <-time.After(time.Second*3):
            fmt.Printf("c1 timeout")
    }

    c2 := make(chan string)
    go func() {
        time.Sleep(time.Second*4)
        c2<-"c2 send"
    }()

    select {
        case msg := <- c2:
            fmt.Printf("Received c2:%s\n", msg)
        case <-time.After(time.Second*3):
            fmt.Printf("c2 timeout")
    }
}


