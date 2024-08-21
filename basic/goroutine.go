package main

import (
    "fmt"
    "time"
)

// Producer function
func producer(jobs chan<- int, done chan<- bool) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Producing job %d\n", i)
        jobs <- i
        time.Sleep(time.Second) // Simulate some work
    }
    close(jobs)
    done <- true
}

// Consumer function
func consumer(jobs <-chan int, done chan<- bool) {
    for job := range jobs {
        fmt.Printf("Consuming job %d\n", job)
        time.Sleep(2 * time.Second) // Simulate some work
    }
    done <- true
}

func main2() {
    jobs := make(chan int, 5)
    done := make(chan bool, 2)

    // Start producer
    go producer(jobs, done)

    // Start consumer
    go consumer(jobs, done)

    // Wait for both producer and consumer to finish
    <-done
    <-done

    fmt.Println("All jobs processed.")
}