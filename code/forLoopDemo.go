package main

import "fmt"

func main(){
	for i := 97; i < 122; i++ {
		fmt.Printf("%c,", i)
	}
	fmt.Printf("%c", 122)
}