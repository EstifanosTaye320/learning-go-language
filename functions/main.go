package main

import (
	"errors"
	"fmt"
)

// function with multiple return vaules
func intdiv(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("dvided by zero")
		return 0, 0, err
	}
	return a / b, a % b, err
}

func main() {
	result, remain, err := intdiv(3, 2)
	switch{ // equivalent to if - else
		case err != nil:
			fmt.Println(err.Error())
		case remain == 0:
			fmt.Println("Result:", result)
		default:
			fmt.Printf("Result: %d, Remainder: %d \n", result, remain)
	}

	var condition = "nice"

	switch condition {// conditional switch
		case "not nice":
			fmt.Println("it was not nice")
		case "nice":
			fmt.Println("it was nice")
		default: 
			fmt.Println("not cool")
	}
}