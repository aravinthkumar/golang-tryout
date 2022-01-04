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
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	// Send only channel syntax
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 52
		wg.Done()
	}(ch, wg)
	wg.Wait()

}
