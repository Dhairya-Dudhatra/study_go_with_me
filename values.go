package main

import (
	"fmt"
)

func main() {
	var hello string
	hello = "demo"
	fmt.Printf("hello value: %s \nhello type: %T,", hello, hello)
	num1, num3 := 2, 5
	fmt.Printf("\nnum1 value: %d \nnum1 type: %T", num1, num1)
	fmt.Printf("\nnum3 value: %d \nnum3 type: %T", num3, num3)

	var bigInt uint64
	//shift-left 64 times and then minus one
	bigInt = 1<<64 - 1
	fmt.Printf("\nbigInt value: %d \nbigINt type: %T", bigInt, bigInt)

	//precision till 6 digits after decimal
	flInt := 5.123456789123456789
	var demo float32
	demo = 4.1234567890123
	fmt.Printf("\nflInt value: %f\nflInt type: %T", flInt, flInt)
	fmt.Printf("\ndemo value: %f\ndemo type: %T", demo, demo)

	var bt []byte
	bt = []byte("hello")
	fmt.Println("\n", bt)
	fmt.Printf("again string version: %s", string(bt))
}
