package main

import "fmt"

func main() {

	// fmt.Println("Hii")
	// fmt.Println("This is from channel")
	// channel()
	// forslelctLoop1()

	// Input
	numbers := []int{2, 3, 4, 5, 6}

	//stage 1
	dataChannel := slicetoChannel(numbers)

	//stage 2
	finalChannel := sq(dataChannel)

	//stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}
