package main

import "fmt"

func main() {
	x := 1

	p := &x
	fmt.Println(p)
	// var p *int = &x   ---> 2nd way-redundant
	// var p *int; p = &x ---> 3rd way-redundant

	fmt.Printf("x=%T &x=%T p=%T *p=%T &p=%T \n\n", x, &x, p, *p, &p)

	fmt.Printf("x=%d &x=%x p=%x *p=%d &p=%x \n", x, &x, p, *p, &p)

	fmt.Println("After dereferencing *p=2")
	*p = 2 // dereferencing
	fmt.Printf("x=%d &x=%x p=%x *p=%d &p=%x \n", x, &x, p, *p, &p)

	fmt.Println("Assigning x=3")
	x = 3
	fmt.Printf("x=%d &x=%x p=%x *p=%d &p=%x \n", x, &x, p, *p, &p)

	fmt.Println("Declaring variable y:=4 and p = &y")
	y := 4
	p = &y
	fmt.Printf("y=%d &y=%x p=%x *p=%d &p=%x \n", y, &y, p, *p, &p)
}
