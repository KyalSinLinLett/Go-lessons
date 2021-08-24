package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	fmt.Println(err)
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read the file - create a buffer of size same as file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)

}
