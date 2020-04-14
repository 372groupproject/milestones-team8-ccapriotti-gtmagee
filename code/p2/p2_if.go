package main

import "fmt"

func main() {
	evaluateBooleanUsingIf(true, true)
	evaluateBooleanUsingIf(true, false)
	evaluateBooleanUsingIf(false, true)
	evaluateBooleanUsingIf(false, false)
}

func evaluateBooleanUsingIf(a bool, b bool){
	if a && b {
		fmt.Println("A is true, B is true")
	} else if a && !b {
		fmt.Println("A is true, B is false")
	} else if !a && b {
		fmt.Println("A is false, B is true")
	}else{
		fmt.Println("A is false, B is false")
	}
}