package main

import "fmt"

func main() {
	for i := 1; i <= 6; i++ {
		switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Everybody in the car, so come on, let's ride")
		}
	}
}