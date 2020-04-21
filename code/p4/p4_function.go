package main

import "fmt"

func main() {
	foo()
	foo(1, 2, 3, 4, 5)

	bar(5)
}

// variadic function
func foo(params ...int) {
	sum := 0

	for i := range params {
		sum += params[i]
	}

	fmt.Printf("sum: %d\n", sum)
}

// good ol' function
func bar(param int) {
	fmt.Printf("sum: %d\n", param)
}
