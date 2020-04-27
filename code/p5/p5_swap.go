package main

import "fmt"

func main() {
	a := 5
	b := 10

	// no temp needed!
	a, b = b, a

	fmt.Printf("a = %d\nb = %d", a, b)

}
