package main

import "fmt"

var x string = "Hello planet"

func main() {
	fmt.Println("main", x)
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

// func f() {
// 	fmt.Println("f", x)
// }
