package array

import (
	"fmt"
	"reflect"
)

func ArrayMain() {
	// simpleArray() // calling simpleArray func
	twoDimensionalArray()
}

func simpleArray() {
	var arr [5]int // array always start from [0,1,2,3,4]
	fmt.Printf("Value of arr %v\n", arr)
	fmt.Printf("Type of arr %T\n", arr)
	fmt.Println(reflect.TypeOf(arr).Kind()) // .Kind returns the specific kind of this type.

	arr[4] = 4 //changing index posting to 4 [0,0,0,4,0]
	//arr[3] = 5
	fmt.Println("setting the value", arr[4])
	fmt.Println("len of arr =", len(arr))

	arr2 := [5]int{4, 5, 6, 2, 8} // this way you can instance intiallize the value
	fmt.Println("value of arr2 =", arr2)
}

func twoDimensionalArray() {
	td := [2][3]int{{1, 2, 5}, {3, 4}} // instance declaration
	fmt.Println(td)
	//td[1][2] = 12
	var twoDArr [2][3]int
	for i := 0; i < len(twoDArr); i++ { // i-> 0,1
		for j := 0; j < len(twoDArr[i]); j++ { // j-> 0,1,2
			twoDArr[i][j] = i + j //twoDArr first loop[0,0], [0,1], [0,2]. sec loop [1,0], [1, 1], [1,2]
		}
	}
	fmt.Println("twoDarr value =", twoDArr)

}
