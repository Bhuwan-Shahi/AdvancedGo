package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg = sync.WaitGroup{}

func doWorld() int {
	return rand.Intn(9999)
}

func main() {
	data := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := doWorld()
				data <- result
			}()
		}
		wg.Wait()
		close(data)
	}()

	for i := range data {
		fmt.Printf("%d\n", i)
	}
}
