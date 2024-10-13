package main

import (
	"fmt"
	"sync"
)

var w = sync.WaitGroup{}

func main() {
	channel := make(chan int)
	for i := 0; i < 5; i++ {
		w.Add(2)
		go func() {
			i := <-channel
			fmt.Println("HI i ma ", i, " from the channel")
			w.Done()
		}()

		go func() {
			channel <- 34
			w.Done()
		}()
	}
	w.Wait()

}
