package values

import "fmt"

//variable usage and decleration
func Variable() {
	var a, b, c = 5, 6, "asdf"
	x := 100

	a += 5 //a = a+5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(x)
	//fmt.Printf("%d", x)
}

// Operator usage
func Values() {
	//Using Arthmetic Operator
	fmt.Println("First" + "Last")

	fmt.Println("10 - 1 = ", 10-1)

}

func Values1() {
	//Using Logical Operator
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)
}
