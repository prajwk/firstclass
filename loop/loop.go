package loop

import "fmt"

func LoopStatememnt() {
	sum := 0
	i := 5
	for i < 10 {
		sum += i //sum = sum + i
		i++      // i = i+1
	}
	fmt.Println(sum)
}

// using nested loop
func GetPrimeNumbers() { // listing prime numbers 2,3,5,7....100
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ { // checking if 2 to 10 are prime numbers are not using nested loop...
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d is prime\n", i)
		}
	}
}

// passing value through func
func GetPrimeNumbers1(last int) { // listing prime numbers 2,3,5,7....100
	var i, j int
	for i = 2; i < last; i++ {
		for j = 2; j <= (i / j); j++ { // checking if 2 to 10 are prime numbers are not using nested loop...
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d is prime\n", i)
		}
	}
}

func Plus() {
	plus, minus := plusMinus(1, 2)
	fmt.Println("a + b =", plus)
	fmt.Println("a + b =", minus)
}

func plusMinus(num1, num2 int) (int, int) {
	plus := num1 + num2
	minus := num1 - num2
	return plus, minus
}

func CheckSwap() {
	var a, b int
	a = 5
	b = 10
	fmt.Printf("Before swapping, a = %d\n", a) //
	fmt.Printf("Before swapping, b = %d\n", b)
	swapByReference(&a, &b) //passing address // call by value case (5, 10) // call by reference case (0x3434, 0x234 )
	fmt.Printf("After swapping, a = %d\n", a)
	fmt.Printf("After swapping, b = %d\n", b)
}

// // call by value
// func swap(x, y int) int {
// 	var temp int
// 	temp = x //save temp to temporary variable
// 	x = y  // put y into x
// 	y = temp // put temp into y
// 	fmt.Printf("Inside swap function swapping, x = %d\n", x)
// 	fmt.Printf("Inside swap function swapping, y = %d\n", y)
// 	return ab

// }

// call by reference
func swapByReference(x, y *int) int { //pointer of integer type
	var temp int
	temp = *x //save temp to temporary variable
	*x = *y   // put y into x
	*y = temp // put temp into y
	// fmt.Printf("Inside swap function swapping, x = %d\n", x)
	// fmt.Printf("Inside swap function swapping, y = %d\n", y)
	return temp
}
