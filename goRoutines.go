package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go sayHello()
		go increaseCounter()
	}
	wg.Wait()

	//WAITGROUPS ARE THERE INORDER TO SYNCRONIZE MULTIPLE GOROUTINES TOGETHER

}

func sayHello() {
	fmt.Println("say hello", counter)
	wg.Done()
}

func increaseCounter() {
	counter++
	wg.Done()
}
