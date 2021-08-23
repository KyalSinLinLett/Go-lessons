package main

import "fmt"

func main() {
	fmt.Print("Enter temperature in F: ")
	var f float64
	fmt.Scanf("%f", &f)
	c := (f - 32.0) * 5 / 9
	fmt.Println("Celsius:", c)
}
