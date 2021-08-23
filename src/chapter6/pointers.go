package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
	fmt.Println(xPtr)  // int pointer - mem addr.
	fmt.Println(*xPtr) // deferenced value inside mem addr. that xPtr points to
}

func main() {
	x := 5
	y := new(int)
	*y = x
	y = &x
	fmt.Println(*y)
	zero(&x) // &x returns the mem addr of x that'll be passed to xPtr
	fmt.Println(&x)
	fmt.Println(x)
}
