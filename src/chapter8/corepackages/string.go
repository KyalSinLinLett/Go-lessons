package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "te"))
	fmt.Println(strings.Index("test", "e"))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))
	fmt.Println(strings.Repeat("a", 5))
	fmt.Println(strings.ToUpper("test"))
}
