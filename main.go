package main

import (
	"fmt"
)

const Pi = 3.14

func main() {
	// fmt.Println("First class")
	// fmt.Printf("The value of Pi is %v", Pi)
	// values.Variable()
	// values.Values()
	// values.Values1()
	// values.StringValue()
	//values.PointersValue()
	// values.IfElseGo()
	//values.SwitchCase()
	//loop.LoopStatememnt()
	//loop.GetPrimeNumbers()
	//loop.GetPrimeNumbers1(100) // passing value through func
	//loop.Plus()
	// loop.CheckSwap()

	// calling instance variable func
	// getSum := func(a, b int) int { // calling func variable inside another func
	// 	return a + b
	// }
	// fmt.Println(getSum(2, 4))
	// fmt.Println(getSum(2, 5))
	// fmt.Println(getSum(2, 10))

	// array.ArrayMain()
	// array.SlicesMain()
	// rangeexample.RangeExample()

	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	fmt.Println("=======")

	nextNumber2 := getSequence()
	fmt.Println(nextNumber2())
	fmt.Println(nextNumber2())
	fmt.Println(nextNumber2())

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
