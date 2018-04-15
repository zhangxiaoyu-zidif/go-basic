package main

import (
	"fmt"
)

func modify(p *int) {
	fmt.Println(p)
	*p = 100
}

func main() {
	var a = 123
	fmt.Println(getParaAddress(&a))

	var p *int
	p = &a
	fmt.Println(*p)

	a++
	fmt.Println(*p)

	modify(&a)
	fmt.Println(*p)
}

func getParaAddress(a *int) *int {
	return a
}
