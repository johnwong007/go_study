package main

import "fmt"
import "time"

func worker(id int, jobs <- chan int, results chan <- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j*2
    }
}

func worker1(id int, jobs1 <- chan int) {
    for j := range jobs1 {
        fmt.Println("worker1", id, "started job", j)
        time.Sleep(time.Second)
        fmt.Println("worker1", id, "finished job", j)
    }
}

func main() {
    jobs := make(chan int, 100)
    jobs1 := make(chan int, 100)
    results := make(chan int, 100)

    for i := 0; i < 3; i++ {
        go worker(i, jobs, results)
        go worker1(i, jobs1)
    }
    
    for i := 0; i < 5; i++ {
        jobs <- i
        jobs1 <- i
    }
    close(jobs)
    close(jobs1)

    for a := 1; a <= 5; a++ {
        <- results
    }
}

