package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.Write([]byte("test"))

	fmt.Println(buf)
}
