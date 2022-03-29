package main

import "fmt"

// function closure
func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	fmt.Println("Functions")
	x := returnFunc("test")
	x()
}
