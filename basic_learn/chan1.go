package main

import "fmt"

func main() {
    messages := make(chan string)

    go func() {messages <- "ping"} ()

    msg := <-messages
    fmt.Println(msg)

    println("简而言之，死锁就是通道有流入没有流出")
    println("死锁的一种演示：")
    messages2 := make(chan string, 2)
    messages2 <- "unused"
    messages2 <- "buffered"
    messages2 <- "string"
    fmt.Println(<-messages2)
    fmt.Println(<-messages2)
}

