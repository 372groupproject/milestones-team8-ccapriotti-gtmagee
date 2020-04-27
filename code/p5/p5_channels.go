package main

import "fmt"

func add(ch chan<- int, m int, n int) {
	ch <- m + n
}

func mul(ch1 <-chan int, ch2 chan<- int, n int) {
	i := <-ch1
	ch2 <- i * n
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go add(ch1, 10, 5)
	go mul(ch1, ch2, 10)

	result := <-ch2 // 150
	fmt.Println("Result:", result)
}
