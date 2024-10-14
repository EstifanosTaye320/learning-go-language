package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var float_32 float32 = 1.34
	var int_32 int32 = 3
	fmt.Println(float_32 + float32(int_32))

	var int1 int = 3
	var int2 int = 2
	fmt.Println(int1 / int2)

	var myString string = "Hello" + " " + "World"
	myString = `Hello
	World`
	fmt.Println(myString)
	fmt.Println(len(myString)) // number of bytes
	fmt.Println(utf8.RuneCountInString(myString)) // Rune = character

	const myConst string = "can't change"
}