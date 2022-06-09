package values

import (
	"fmt"
	"strings"
)

func StringValue() {
	fmt.Println(strings.Compare("a", "b")) //-1
	fmt.Println(strings.Compare("a", "a")) //0
	fmt.Println(strings.Compare("b", "a")) //1
}
