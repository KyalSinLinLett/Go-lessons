package main

import "fmt"

func main() {
	xs := []int{23, 234, 45, 46, 646}

	fmt.Println(add(xs...))
	fmt.Println(add(1, 2, 3, 4, 5, 6))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}

	return total
}
