package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int) // rtrn a pointer to int
	one(xPtr)
	fmt.Println(*xPtr)
}
