package main

import "fmt"

func setN(n *int, val int) {
	*n = val
}

func main() {
	n := 0
	setN(&n, 10)
	fmt.Printf("Value of n is %d", n)
}
