package main

import (
	"fmt"
)

func main() {
	// fmt.Println(add(2, 21))
	// fmt.Println(swap1("World", "Hello"))
	// fmt.Println(swap2("World", "Hello"))
	scores1 := []float32{91, 82, 99}
	fmt.Printf("average: %.2f\n", avg(scores1))

	scores2 := []float32{72, 81, 78, 91, 68}
	fmt.Printf("average: %.2f\n", avg(scores2))

	employees := map[string]map[string]string{
		"BT": {
			"firstName": "John",
			"lastName":  "Doe",
		},
		"PC": {
			"firstName": "John",
			"lastName":  "D",
		},
	}
	if emp, ok := employees["BT"]; ok {
		fmt.Println(emp["firstName"], emp["lastName"])
	}

	for initials, emp := range employees {
		fmt.Println(initials, emp["firstName"], emp["lastName"])
	}

}

func add(a, b int) int {
	return a + b
}

func swap1(x, y string) (string, string) {
	return y, x
}

func swap2(x, y string) (r1, r2 string) {
	r1 = y
	r2 = x
	// since we are returning named values we don't need to include them in the
	// return statement. Just pass return
	return
}

func avg(scores []float32) float32 {
	var total float32
	for _, score := range scores {
		total += score
	}

	return total / float32(len(scores))
}
