package main

import "fmt"

type Gator struct {
	teeth int
	age int
	name string
}

func main(){
	var myGator Gator = Gator{100, 10, "Jeremy"}
	fmt.Printf("Gator Teeth Count: %d\n", myGator.teeth)
	fmt.Printf("Gator Age: %d\n", myGator.age)
	fmt.Printf(myGator.name + " is Hungry...")
}