package main

import "fmt"

type Point struct {
	x int32
	y int32
}

// embedded struct
type Circle struct {
	radius float64
	center *Point
}

func main() {
	fmt.Println("Struct Types in go")
	// var p1 Point = Point{1, 2, true}
	// var p2 Point = Point{-5, -7, false}
	// p1.x = 10
	// fmt.Println(p1.x, p2.x)
	c1 := Circle{4.56, &Point{4, 5}}
	fmt.Println(c1.center.x)
	c1.center.y = 12
	fmt.Println("Circle: ", c1.center, c1.center.x, c1.center.y)

}
