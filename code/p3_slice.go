package main

import "fmt"
import "strconv"

func main(){
	testSlice := []int{5,1,2,3,10}
	fmt.Println("OG Value: " + strconv.Itoa(testSlice[len(testSlice) - 1]))
	subSlice := testSlice[2:]
	subSlice[2] = -99
	fmt.Println("Modified in New Slice: " + strconv.Itoa(subSlice[2]))
	fmt.Println("Modified in Old Slice: " + strconv.Itoa(testSlice[4]))
	fmt.Println("Length of Slice: " + strconv.Itoa(len(testSlice)))
	testSlice = append(testSlice, 501)
	fmt.Println("Apppended Val: " + strconv.Itoa(testSlice[5]))
}