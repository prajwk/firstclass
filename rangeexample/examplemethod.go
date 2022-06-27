package rangeexample

import (
	"fmt"
)

type Rect struct {
	width, height int
}

// using method function
func (r *Rect) area() int { // defining struct infront of func area()
	return r.width * r.height
}

func (r Rect) perim() int { // In Go it automatically handles even if we dont use pointer
	return 2*r.width + 2*r.height
}

func MethodExample() {
	rectangle := Rect{width: 100, height: 20}
	fmt.Println("Area", rectangle.area())
	fmt.Println("Perimeter", rectangle.perim())

	// by passing address using "&rectangle"
	r := &rectangle
	fmt.Println("Area", r.area())
	fmt.Println("Perimeter", r.perim())
}
