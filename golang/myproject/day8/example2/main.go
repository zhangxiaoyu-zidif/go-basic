package main

import (
	"fmt"
)

type add_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func operator(op add_func, a, b int) int {
	return op(a, b)
}

func main() {
	c := add
	fmt.Println(c)
	fmt.Println(operator(c, 1, 2))
}
