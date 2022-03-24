package main

import (
	"fmt"
)

func main() {
	fmt.Println("For and while loops..")
	// x := 0
	// for x <= 5 {
	// 	fmt.Println("Current iteration: ", x)
	// 	x++
	// }

	for x := 0; x <= 5; x++ {
		if x == 3 {
			continue
		}
		fmt.Println(x)
	}
}
