package main

import "fmt"

func main() {
	bytes := []byte("Xpo(u!zpv!ublf!nf!up!gvolz!upxo@")

	for _, b := range bytes {
		fmt.Printf("%c", b-1)
	}
}
