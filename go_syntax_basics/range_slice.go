package main

import (
	"fmt"
)

func main() {
	fmt.Println("Range and Slice Examples")
	var a []int = []int{1, 3, 4, 56, 7, 12, 4, 6}

	// traditional for loop
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// using range
	for _, element := range a { // iterating through index and element
		// The underscore is an anonymous variable - you can not use it
		fmt.Printf("Element %d\n", element)
	}
}
