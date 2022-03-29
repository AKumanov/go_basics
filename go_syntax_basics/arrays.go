package main

import (
	"fmt"
)

func main() {
	arr := [...]int{4, 5, 6, 7}
	var sum int
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		sum += arr[i]
	}

	fmt.Printf("Sum of the array: %v", sum)

	arr2D := [...][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	fmt.Println(arr2D)

	for i := 0; i < len(arr2D); i++ {
		for j := 0; j < len(arr2D[i]); j++ {
			fmt.Printf("Column %v and row %v \n", arr2D[i][j], arr2D[j][i])
		}
	}

}
