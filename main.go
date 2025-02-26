package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		defer close(ch)
		ch <- 10
	}()

	msg := <-ch
	fmt.Println(msg)
}
