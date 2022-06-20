package array

import (
	"fmt"
)

func SlicesMain() {
	intro()
	// apendValue()
}

func apendValue() {
	s := make([]string, 2)
	fmt.Println("slice", s)
	fmt.Printf("len = %v, cap = %v\n", len(s), cap(s))

	s[0] = "g"
	s[1] = "o"
	fmt.Println("slices =", s)
	fmt.Printf("len = %v, cap = %v\n", len(s), cap(s))

	// append is use to add value in slice
	s = append(s, "l")
	fmt.Println("slices =", s)
	fmt.Printf("len = %v, cap = %v\n", len(s), cap(s))

}

func intro() {
	// var arr = []int{1,3,5}
	// fmt.Printf("Value of arr %v\n", arr)
	// fmt.Printf("Type of arr %T\n", arr)
	// fmt.Printf("length of arr %v\n", len(arr))
	// fmt.Println(reflect.TypeOf(arr).Kind()) // .Kind returns the specific kind of this type.

	s := make([]string, 6) // make(type, len , cap)
	fmt.Println("slices:", s)

	s[0] = "g"
	s[1] = "o"
	s[2] = "l"
	s[3] = "a"
	s[4] = "n"
	s[5] = "g"
	fmt.Println("slices:", s)
	fmt.Printf("len = %v, cap = %v\n", len(s), cap(s))

	// using slicing operation.
	s = s[:]
	fmt.Println(s)
	fmt.Printf("case1 len = %v, cap = %v\n", len(s), cap(s))

	s = s[2:4] // 2 is included 4 is not included
	fmt.Println(s)
	fmt.Printf("case2 len = %v, cap = %v\n", len(s), cap(s))

	s = s[:4]
	fmt.Println(s)
	fmt.Printf("case3 len = %v, cap = %v\n", len(s), cap(s))

	s = s[2:]
	fmt.Println(s)
	fmt.Printf("case4 len = %v, cap = %v\n", len(s), cap(s))

	s = s[:]
	fmt.Println(s)
	fmt.Printf("case5 len = %v, cap = %v\n", len(s), cap(s))

}
