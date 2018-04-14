package main

import (
	"fmt"
)

func swap(a, b int) (int, int) {
	return b, a
}
func main() {
	var a, b int = 3, 4
	fmt.Println(swap(a, b))
}
