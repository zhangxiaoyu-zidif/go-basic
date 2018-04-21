package main

import "fmt"

func add(a int, arg ...int) int {
	var sum = a
	defer fmt.Println(sum)
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}

func concat(s string, arg ...string) string {
	for _, str := range arg {
		s += str
	}
	return s
}

func main() {
	sum := add(10, 20, 30)
	fmt.Println(sum)
	sum2 := add(1, 2, 3)
	fmt.Println(sum2)

	unitString := concat("abc", "123", "xyz")
	fmt.Println(unitString)
}
