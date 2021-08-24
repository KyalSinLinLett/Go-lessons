package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("custom error message")
	fmt.Println(err)
}
