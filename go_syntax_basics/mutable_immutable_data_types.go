package main

import "fmt"

func main() {
	fmt.Println("Mutable vs. immutable data types")
	// var x int = 5
	// y := x
	// y = 7
	// fmt.Println(x, y)

	var x []int = []int{3, 4, 5}
	y := x
	y[0] = 100
	fmt.Println(x, y)

}
