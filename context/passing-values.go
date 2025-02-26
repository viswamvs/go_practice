package main

import (
	"context"
	"fmt"
)

// Function that retrieves user ID from context
func checkAuth(ctx context.Context) {
	userID := ctx.Value("userID") // Retrieve value
	if userID != nil {
		fmt.Println("Authenticated user:", userID)
	} else {
		fmt.Println("No user found in context")
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "userID", 42) // Store user ID

	checkAuth(ctx) // Access user ID in another function
}
