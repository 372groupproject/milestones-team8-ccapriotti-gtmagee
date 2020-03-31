package main

import "fmt"

func main(){
	for i := 97; i < 122; i++ {
		fmt.Printf("%c,", i)
	}
	fmt.Printf("%c\n", 122)
	
	var myArray [4]string = [4]string{"this", "land", "is", "forsaken"}
	for i, myStr := range myArray {
		fmt.Println(myStr,i)
	}
}