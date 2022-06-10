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
