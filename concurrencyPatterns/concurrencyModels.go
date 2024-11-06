package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func printNum(k int) {
	fmt.Println(k)
}

// channels are fifi ques
func channel() {
	myChannel := make(chan string)
	Channel1 := make(chan string)

	go func() {
		//this is join point in the fork join model
		myChannel <- "bhuwan"
	}()
	go func() {
		//this is join point in the fork join model
		Channel1 <- "bhuwan i"
	}()
	select {
	case message := <-myChannel:
		fmt.Println(message)
	case message1 := <-Channel1:
		fmt.Println(message1)
	}
}
