package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task completed successfully!")
	case <-ctx.Done():
		fmt.Println("Task cancelled due to deadline:", ctx.Err())
	}
}

func deadline() {
	deadline := time.Now().Add(2 * time.Second) // Set a fixed deadline
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go task(ctx)

	time.Sleep(3 * time.Second) // Allow time to observe output
}
