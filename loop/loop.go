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

// func swap(x, y string) (string, string) {
// 	return y, x
// }
