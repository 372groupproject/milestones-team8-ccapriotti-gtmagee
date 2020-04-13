package main

import "fmt"

func main(){
	signed()
	unsigned()

	var a int = 5
	var b int = 6
	fmt.Println("a + b =", a + b)
}

func signed(){
	var sevenOnes int8 = 0xff >> 1
	var fifteenOnes int16 = 0xffff >> 1
	var thirtyOneOnes int32 = 0xffffffff >> 1
	var sixtyThreeOnes int64 = 0xffffffffffffffff >> 1
	fmt.Printf("Max Num 8 Bits Signed: %d\n", sevenOnes);
	fmt.Printf("Max Num 16 Bits Signed: %d\n", fifteenOnes);
	fmt.Printf("Max Num 32 Bits Signed: %d\n", thirtyOneOnes);
	fmt.Printf("Max Num 64 Bits Signed: %d\n\n", sixtyThreeOnes);
}

func unsigned(){
	var eightOnes uint8 = 0xff
	var sixteenOnes uint16 = 0xffff
	var thirtyTwoOnes uint32 = 0xffffffff
	var sixtyFourOnes uint64 = 0xffffffffffffffff
	fmt.Printf("Max Num 8 Bits Unsigned: %d\n", eightOnes);
	fmt.Printf("Max Num 16 Bits Unsigned: %d\n", sixteenOnes);
	fmt.Printf("Max Num 32 Bits Unsigned: %d\n", thirtyTwoOnes);
	fmt.Printf("Max Num 64 Bits Unsigned: %d\n\n", sixtyFourOnes);
}