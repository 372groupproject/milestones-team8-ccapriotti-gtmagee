package main

import (
	"fmt"
	"time"
)

func foo(n int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(n)
	}
}

func main() {
	// start goroutine
	go foo(1)
	foo(2)
}
