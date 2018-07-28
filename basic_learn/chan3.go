package main

import "fmt"

func ping(chan1 chan<- string, msg string) {
    chan1<- msg
}

func pong(chan1 <-chan string, chan2 chan<- string) {
    msg := <- chan1
    chan2 <- msg
}

func main() {
    chan1 := make(chan string, 1)
    chan2 := make(chan string, 1)
    ping(chan1, "hello")
    pong(chan1, chan2)
    
    msg := <- chan2
    fmt.Println(msg)
    
}

