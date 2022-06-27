package rangeexample

import "fmt"

func RangeExample() {
	person1 := person{
		firstName:   "prajwal",
		lastName:    "shrestha",
		phoneNumber: "2345342352",
		age:         28}

	if person1.age > 30 {
		fmt.Println("age = ", person1.firstName)
	} else {
		fmt.Println("Person =", person1.lastName)
	}
	fmt.Println("Person1 =", person1)

	person2 := person{
		firstName: "sami",
		lastName:  "shrestha"}
	fmt.Println("Person2 =", person2)

	fmt.Println(newPerson("ram", "gautam"))

	// exampleOne()
	// exampleMap()
	// exampleTwo()
	// exampleVariadicFunction()
}



type person struct {
	/* struct is a collection of different variable  */
	firstName   string
	lastName    string
	phoneNumber string
	age         int
}

func newPerson(firstName string, lastName string) *person { //using pointer to pass address
	p := person{
		firstName: firstName,
		lastName:  lastName,
	}
	p.age = 29
	return &p // passing address of p using "&p"
}

func exampleVariadicFunction() {
	nums := []int{2, 3, 4, 5, 6} // passing array
	// sum(1, 2, 3, 4, 5, 6) // passing multiple arguments case
	sum(nums...) // triple dots will pass array as multiple arguments
	sum(1, 2)
	sum(1, 2, 4, 7)
}

func sum(nums ...int) {
	fmt.Println("nums value = ", nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum =", sum)
}

func exampleTwo() {
	for i, value := range "ABCabc" { //using ascii value
		fmt.Println(i, value)
	}
}

func exampleMap() {
	// make(map[key_type]val+type) Map's key_type is always unique, it cannot be different.
	n := map[string]int{"one": 1, "two": 2}
	fmt.Println("map = ", n)

	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	delete(m, "three")
	fmt.Println("map = ", m)

	/* In map range gives two return type Key and value also has two output
	range => (key, value)
	*/
	for key, value := range m {
		fmt.Printf("%v = %v\n", key, value)
	}
}

func exampleOne() {
	nums := []int{2, 3, 4, 5}
	sum := 0
	// sum = nums[0] + nums[1] + nums[2] + nums[3]
	for i := 0; i < len(nums); i++ {
		sum += nums[i] // sum = sum + nums[i]
	}
	fmt.Println("sum = ", sum)

	/* range gives two return type index and value also has two output
	range => (index, value)
	*/

	sum = 0
	for _, num := range nums { // underscore '_' used for unused variable
		// fmt.Println(i)
		sum += num
	}
	fmt.Println("sum = ", sum)
}
