// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// var wg = sync.WaitGroup{}
// var counter = 0

// // this is readwrite mutex that allows you to give to lock and objce and unlock to prevent from the race condiitons
// var m = sync.RWMutex{}

// // ,utex is there inorder to give the lock and unlock to prevent the race condition
// func main() {
// 	runtime.GOMAXPROCS(100)
// 	for i := 1; i < 10; i++ {
// 		wg.Add(2)
// 		m.RLock()
// 		go sayHello()
// 		m.Lock()
// 		go increaseCounter()
// 	}
// 	wg.Wait()

// 	//WAITGROUPS ARE THERE INORDER TO SYNCRONIZE MULTIPLE GOROUTINES TOGETHER

// }

// func sayHello() {
// 	fmt.Println("say hello", counter)
// 	m.RUnlock()
// 	wg.Done()
// }

// func increaseCounter() {
// 	counter++
// 	m.Unlock()
// 	wg.Done()
// }
