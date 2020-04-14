package main

import "fmt"

func main(){
	fmt.Println("Main: 1")
	a()
	fmt.Println("Main: 2")
}

func a(){
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Saved in A", r)
        }
    }()
	fmt.Println("A: 1")
	b()
	fmt.Println("A: 2")
}

func b(){
	fmt.Println("B: 1")
	c()
	fmt.Println("B: 2")
}

func c(){
	fmt.Println("C: 1")
	panic("PANICKED IN C")
	fmt.Println("C: 2")
}