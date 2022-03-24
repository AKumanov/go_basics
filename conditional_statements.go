package main

import (
	"fmt"
)

/*
-----------VARIABLES-------------------

---strings
---int
---uint - unsigned integer - any positive whole number
---float64 - decimal value
---& - POINTER - points to the memory location of a variable - allows access to it
-----------ENDVARIABLES----------------
*/

func main() {
	fmt.Println("Welcome to the test quiz")
	var name string

	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)
	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("You are old enough!")
	} else {
		fmt.Println("You are not old enough!")
		return
	}
	fmt.Printf("only if expression evaluates to true..")

}
