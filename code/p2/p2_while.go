package main

import "fmt"

func main(){
	i := 0
	target := 420
	myArr := [5]int{3,90,69,47,420}
	for i != -1 && myArr[i] != target {
		i++
		if i == len(myArr){
			i = -1
		}
	}
	fmt.Printf("Target: %d, Index: %d\n", target, i)
}