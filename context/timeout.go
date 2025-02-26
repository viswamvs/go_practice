package main

import (
	"context"
	"fmt"
	"time"
)

func slowTask(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second): // Simulating long task
		fmt.Println("Task completed!")
	case <-ctx.Done(): // If timeout occurs, exit
		fmt.Println("Task cancelled due to timeout:", ctx.Err())
	}
}

func timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Ensure cleanup

	go slowTask(ctx) // Start task in goroutine

	time.Sleep(3 * time.Second) // Wait before program exits
}
