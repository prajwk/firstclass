package iexample

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	width, height float64
}

// using method function
func (r *Rect) area() float64 { // defining struct infront of func area()
	return r.width * r.height
}

func (r Rect) perim() float64 { // In Go it automatically handles even if we dont use pointer
	return 2*r.width + 2*r.height
}

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius // using math func --> math.Pi
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//using interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(int64(g.area())) // type conversion adding int64
	fmt.Println(g.perim())
}

func MainIStart() {
	rectangle := Rect{width: 10, height: 5}
	circle := Circle{radius: 5}

	fmt.Println("Area of rectangle =", rectangle.area())
	fmt.Println("Perimeter of rectangle =", rectangle.perim())
	fmt.Println("Area of Circle =", circle.area())
	fmt.Println("Perimeter of Circle =", circle.perim())

	measure(&rectangle)
	measure(&circle)
}
