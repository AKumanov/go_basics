package main

import "fmt"

/*
& - get pointer
* - dereference
*/

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	fmt.Println("Pointers in go")
	// x := 7
	// // look at the reference
	// fmt.Println(&x) // this is a pointer to the place in the memory, where the variable is saved
	// y := &x         // y is equall to the pointer of x
	// fmt.Println(y)  // prints 0xc0000180a8
	// *y = 8          // dereference - take this pointer value and dereference it
	// fmt.Println(x)  // now x is changed as well
	toChange := "hello!"
	// fmt.Println(toCHange)
	// changeValue(&toCHange)
	// fmt.Println(toCHange)
	var pointer *string = &toChange
	fmt.Println(&pointer)
}
