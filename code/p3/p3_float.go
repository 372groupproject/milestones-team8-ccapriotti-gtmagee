package main

import "fmt"

func main() {
	var pi float32 = 3.14159
	var r int = 5

	// must cast int to float to perform arithmetic
	fmt.Printf("Area of circle: %f\n", (pi * (float32(r * r))))

	fmt.Printf("Float divided by float: %f\n", (3.0 / 10.0))
	fmt.Printf("Float divided by int: %f\n", (3.0 / 10))
	fmt.Printf("Int divided by float: %f\n", (3 / 10.0))
}
