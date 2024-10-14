package main

import "fmt"

func main() {
	var intArr [3]int
	fixedimplied:= [...]int{4,5,6}
	myslice := []int{1,2,3}
	myslice3 := []int{4,5,6}

	fmt.Println(fixedimplied[0])
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3]) // sliceing an array

	// MEMORY LOCATIONS
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
	
	myslice = append(myslice, myslice3...) // spread operator
	fmt.Println(myslice)

	myMap:= map[string]int{"adam":4}
	var value, isinmap = myMap["estif"]
	if isinmap == true {
		fmt.Println(value)
	} else {
		fmt.Println("estif is not a key in myMap")
	}

	delete(myMap, "adam") // no return value

	// while loop with for
	var i = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}