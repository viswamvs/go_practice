package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random number generator

	jobs := make(chan int, 5) // Buffered channel
	var wg sync.WaitGroup

	// Start Producer
	go producer(jobs, 10)

	// Start Consumers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go consumer(i, jobs, &wg)
	}

	wg.Wait() // Wait for all consumers to finish
	fmt.Println("All jobs processed.")
}

func producer(jobs chan<- int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("Produced %d\n", i)
		jobs <- i
	}
	close(jobs)
}

func consumer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Consumer %d processed %d\n", id, job)
	}
}
