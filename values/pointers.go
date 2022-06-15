package values

import "fmt"

func PointersValue() { // address and reference
	var p *int

	i := 42
	p = &i
	fmt.Println(*p)
}
