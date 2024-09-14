package main

import "fmt"

func main() {
	defer fmt.Println("Go Defer") // follows LIFO order
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
	fmt.Println("Welcome to")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

/*A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns,
either because the surrounding function executed a return statement , reached the end of its function body,
or because the corresponding goroutine is panicking .*/
