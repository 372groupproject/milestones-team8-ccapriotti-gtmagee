package main

import "fmt"

func main(){
	defer fmt.Println("I will be printed second")
	fmt.Println("I will be printed first")
}