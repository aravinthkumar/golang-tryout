package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		for msg := range ch { // Would raise a deadlock hence closing of channel is required
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
		close(ch)
	}(ch, wg)
	wg.Wait()

}
