package main

import "fmt"

func main() {
	xs := []float64{93, 98, 77, 82, 83}
	avg := average(xs)

	fmt.Println(avg)
}

func average(xs []float64) float64 {
	total := 0.0

	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}
