package main

import (
	"context"
	"fmt"
	"time"
)

func cancellation() {
	context, cancel := context.WithCancel(context.Background())

	go worker(context, 1)
	time.Sleep(3 * time.Second)

	cancel()

	fmt.Println("cancelling the worker...")
	time.Sleep(time.Second)
}

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done(): // If context is cancelled, exit the goroutine
			fmt.Printf("Worker %d stopping...\n", id)
			return
		default:
			fmt.Printf("Worker %d is working...\n", id)
			time.Sleep(time.Second) // Simulate work
		}
	}
}
