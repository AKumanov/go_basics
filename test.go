package main

import (
	"fmt"
)

func calculation(word string) (result int) {
	if len(word)%2 == 0 {
		result += len(word)
	}
	return result
}

func main() {
	final := 0
	myArr := [...]string{
		"hello",
		"world"}

	fmt.Println(myArr)
	for index := range myArr {
		final += calculation(myArr[index])
	}
	fmt.Println(final)
}
