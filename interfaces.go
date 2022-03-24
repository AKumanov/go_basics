package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Interfaces in go..")
	c1 := circle{radius: 4.5}
	fmt.Println(c1.area())
	r1 := rectangle{width: 3.8, height: 4.3}
	fmt.Println(r1.area())
	shapes := []shape{c1, r1}
	var shape1 float64 = shapes[0].area()
	var shape2 float64 = shapes[1].area()
	fmt.Println("SHAPE1, SHAPE2 ", shape1, shape2)

	for _, shape := range shapes {
		fmt.Printf("Shape %v has an area of %f\n", shape, getArea(shape))
	}

}
