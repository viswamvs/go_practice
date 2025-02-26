package main

import (
	"fmt"
	"sync"
)

func main() {
	workers := 3
	tasks := 10

	var wg sync.WaitGroup

	jobs := make(chan int, tasks)
	results := make(chan int, tasks)

	// sending value to a input channel
	for i := 1; i <= tasks; i++ {
		jobs <- i
	}

	close(jobs) // close the input channel after sending the data

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	wg.Wait()
	close(results) // close the output channel after receiving the data

	for v := range results {
		fmt.Println(v)
	}
}

func worker(id int, input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range input {
		fmt.Printf("Worker %d Processed %d Job\n", id, v)
		output <- v * v
	}
}
