package main

import "fmt"

func main() {
	// x := make(map[string]int)

	// x["key"] = 10

	// val, ok := x["key"]

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
	}

	fmt.Println(elements)
	fmt.Println(elements["H"])
	fmt.Println(elements["H"]["name"])
	fmt.Println(elements["H"]["state"])

}
