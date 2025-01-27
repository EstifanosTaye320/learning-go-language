package main

import "fmt"

func main() {
	age := 23

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is not less than 40")
	}

	names := []string{"ted", "fred", "free", "lance"}

	for index, name := range names {
		if index == 1 {
			fmt.Println("Continuing at index", index)
			continue
		}

		if index > 2 {
			fmt.Println("Breaking at index", index)
			break
		}

		fmt.Println(name)
	}
}
