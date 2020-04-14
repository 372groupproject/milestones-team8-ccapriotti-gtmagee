package main

import "fmt"

func main() {
	var a int = 1
	var b int = 2

	var x bool = a == b
	var y bool = a != b
	var z bool = (x && y) || true

	fmt.Printf("x = %v\n", x)
	fmt.Printf("y = %v\n", y)
	fmt.Printf("z = %v\n", z)
}
