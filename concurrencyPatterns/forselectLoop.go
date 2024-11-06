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

func forslelctLoop1() {
	go func() {
		for {
			select {
			default:
				fmt.Println("Doing work")
			}
		}
	}()
	time.Sleep(time.Second * 4)
}
