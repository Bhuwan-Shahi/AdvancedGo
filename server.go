package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {

// 	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("helo wev dev!"))
// 		fmt.Println("Hello world")
// 	})
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	a := func() {
// 		fmt.Println("hello world for the anonymous function")
// 	}
// 	fmt.Print(a)

// }
