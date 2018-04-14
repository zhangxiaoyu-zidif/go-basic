package main

import (
	"fmt"
)

func modify(a *int32) {
	*a = *a + 1
	return
}

func minus(a *int32) {
	*a = *a - 1
	return
}

func main() {
	var a int32 = 10
	modify(&a)
	fmt.Println(a)

	minus(&a)
	fmt.Println(a)
}
