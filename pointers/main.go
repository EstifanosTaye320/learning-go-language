package main

import "fmt"

func main() {
    var variablie = 4
    fmt.Println(&variablie)

    var new = variablie
    
    fmt.Println(new)
}