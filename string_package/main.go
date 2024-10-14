package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := []rune("resume")

	for i:= range myString {
		fmt.Printf("index: %v, value: %v", i, myString[i])
	} // int32 version of the characters


	var strSlice = []string{"s", "n", "e", "a", "k"}
	var strBuilder strings.Builder

	for i:= range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Println(catStr)
}