package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Go..")
	var mp map[string]int = map[string]int{
		"apples": 5,
		"pear":   6,
		"orange": 9,
	}

	fmt.Println(mp)

	new := make(map[string]int)
	fmt.Println(new)

	// check if a key exists
	val, ok := mp["test"] // val stores the variable, ok stores a bool
	fmt.Println(val, ok)

	// iterate through maps in go
	for key, value := range mp {
		fmt.Println(key, value)
	}

}
