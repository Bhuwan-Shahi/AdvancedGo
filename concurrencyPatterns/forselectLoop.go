package main

import (
	"fmt"
	"time"
)

func forslelctLoop() {

	//Buffered channels because of 3 in the last
	charChannel := make(chan string, 3)
	chars := []string{"hello", "how", "are"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}
	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}
}

// done channel
func dowork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing work")
		}
	}
}

//For select loop pattern is used when parent want to close a goroutine using channels

func forslelctLoop1() {
	done := make(chan bool)

	dowork(done)
	time.Sleep(time.Second * 2)
	close(done)
}
