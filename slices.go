package main

import (
	"fmt"
)

/*
--------------SLICES--------------
*/
func main() {
	fmt.Println("Slices tutorial")
	var x [5]int = [5]int{1, 2, 3, 4, 5} //array
	fmt.Printf("%T", x)
	var s []int = x[1:3] //slice
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))     // capacity - because we can only take the items from the right of the start of the slice
	fmt.Println(s[:cap(s)]) // takes the most available items starting from index 1

	// declaration of a slice
	var a []int = []int{5, 6, 7, 8, 9}
	fmt.Println(a)

	// adding elements to slice
	a = append(a, 10)
	fmt.Println(a)

	// create slices with make

	b := make([]int, 5) // in the params, first we declare the type of object, then the capacity of it
	fmt.Println(b)
	b[0] = 10
	fmt.Println(b)

	// checking the type of the object
	fmt.Printf("%T", b)

}
