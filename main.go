package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// By default the channels are bi-directional
	ch := make(chan int)

	wg.Add(2)
	// Receive Only channel syntax
	go func(ch <-chan int, wg *sync.WaitGroup) {
		if msg, ok := <-ch; ok {
			fmt.Print(msg, ok)
		} // Wouldn't print at all since the channel is closed
		wg.Done()
	}(ch, wg)
	// Send only channel syntax
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// Only Send only channels can close
		close(ch)
		wg.Done()
	}(ch, wg)
	wg.Wait()

}
