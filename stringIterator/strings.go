package main

import "fmt"

func main() {
	// String with each of ascii characters (from 0 to 127)
	asciiStr := "Hello, I am Dhairya. How are you ?"

	// String with multiple byte (non-ascii) characters (unicode)
	// Examples: emoji, Chinese, accented, Arabic, Hindi etc.
	// multiByteStr := "你好, мир, café, 😀, عربى, हिन्दी"

	fmt.Println(asciiStr[1:3] == "el")

	// for i := 0; i < len(asciiStr); i++ {
	// 	fmt.Printf("%c", asciiStr[i])
	// }
	// fmt.Println()
	// for i := 0; i < len(multiByteStr); i++ {
	// 	fmt.Printf("%c", multiByteStr[i])
	// }

	// rune1 := []rune(asciiStr)
	// rune2 := []rune(multiByteStr)

	// // for i := 0; i < len(rune1); i++ {
	// // 	fmt.Printf("%c", rune1[i])
	// // }
	// // fmt.Println()
	// // for i := 0; i < len(rune2); i++ {
	// // 	fmt.Printf("%c", rune2[i])
	// // }
	// fmt.Println()
	// fmt.Println()
	// for ind, char := range rune1 {
	// 	fmt.Printf("Index: %d, Character: %b\n", ind, char)
	// }
	// for ind, char := range rune2 {
	// 	fmt.Printf("Index: %d, Character: %b\n", ind, char)
	// }

}
