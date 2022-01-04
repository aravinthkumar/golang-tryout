package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// Buffered Channel with value 1
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		// This is ignored since buffer is 1
		ch <- 49
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	wg.Wait()

}
