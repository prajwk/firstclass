package values

import "fmt"

func IfElseGo() {
	// var x int
	// fmt.Print("Enter the x = ")
	// fmt.Scanln(&x)

	if x := 10; x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}
}

func SwitchCase() {
	x := 100
	switch {
	case x > 100:
		fmt.Println("case 0")
	case x > 10:
		fmt.Println("case 10")
	default:
		fmt.Println("default")
	}
}
