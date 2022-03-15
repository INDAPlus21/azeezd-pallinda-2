package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() { // Goroutine here makes the program able to listen in line 11
		ch <- "Hello world!"
	}()
	fmt.Println(<-ch)
}
