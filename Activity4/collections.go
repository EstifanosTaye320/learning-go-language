package main

import "fmt"

func main() {
	arr:= [3]int{1, 2, 3}
	fmt.Println("Array", arr)
	fmt.Println("arr", len(arr))
	fmt.Println("arr", cap(arr))

	slice:= []int{1, 2, 3}
	slice  = append(slice, 7)
	fmt.Println("Slice", len(slice))
	fmt.Println("Slice", cap(slice))

	myMap:= make(map[string]int)
	myMap["alice"] = 25
	myMap["Bob"] = 30
	fmt.Println("Map: ", myMap)
	fmt.Println("Alice's age is ", myMap["alice"])

	for i, v:= range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	for key, value := range myMap {
		fmt.Printf("%s is %d years old \n", key, value)
	}
}