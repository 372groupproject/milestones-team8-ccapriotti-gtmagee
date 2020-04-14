package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["key"] = 1

	fmt.Println("map:", m)

	val := m["key"]
	fmt.Println("value: ", val)
}
