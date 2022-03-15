package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // Global to be accessed by all goroutines

func main() {
	ch := make(chan int)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		wg.Add(1) // There is something to wait on
		ch <- i
	}

	wg.Wait() // Wait until all prints are done
	close(ch)
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
		wg.Done() // Delay and printing is done!
	}
}
