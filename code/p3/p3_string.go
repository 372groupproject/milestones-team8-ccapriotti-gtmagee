package main

import "fmt"
import "strings"

func main(){
	var myString string = "corona is bad, yes?"
	fmt.Println("myString = 'corona is bad, yes?'")
	fmt.Println("myString contains 'bad' =", strings.Contains(myString, "bad"))
	fmt.Println("Concatenate myString + ' duh':", myString + " duh.")
	fmt.Printf("Second Char of myString: %c\n", myString[1])
	fmt.Println("myString From Chars 11 to 13: '" + myString[10:13] + "'")
	fmt.Println("Index of 'i' in myString:", strings.Index(myString, "i"))
	fmt.Println("myString in caps: " + strings.ToUpper(myString))
	fmt.Println("myString Split: [" + strings.Join(strings.Split(myString, ","), "#") + "]")
}