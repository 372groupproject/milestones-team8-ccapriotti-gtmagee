package main

import "fmt"
import "strconv"

func main(){
	var myArr [2]string = [2]string{"howdy", "partner"}
	myArr[0] = "hello"
	fmt.Println("Element 0: " + myArr[0])
	fmt.Println("Element 1: " + myArr[1])
	fmt.Println("Length: " +  strconv.Itoa(len(myArr)))
}