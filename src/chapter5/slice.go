package main

import "fmt"

func main() {
	slice1 := []float64{2, 3, 4}
	slice2 := append(slice1, 5, 6, 7)

	fmt.Println(slice1)
	fmt.Println(slice2)

	sliceA := []float64{1, 2, 3}
	sliceB := make([]float64, 2)
	copy(sliceB, sliceA)

	fmt.Println(sliceA)
	fmt.Println(sliceB)
}
