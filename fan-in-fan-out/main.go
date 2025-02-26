package main

import (
	"fmt"
	"sync"
)

func main() {

	nums := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup

	input := make(chan int, len(nums))
	output := make(chan int, len(nums))

	// sending values to input channels
	for _, v := range nums {
		input <- v
	}

	close(input) // Close input channel after sending all numbers

	// Fan out (Starting multiple workers)
	numberOfWorkers := 2
	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go worker(input, output, &wg)
	}

	wg.Wait()
	close(output) // Close output channel after all workers finish

	// receiving values from a channel
	for v := range output {
		fmt.Println(v)
	}
}

func worker(input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range input { // receiving values from input channel
		output <- v * v // sending values to a output channel
	}
}
