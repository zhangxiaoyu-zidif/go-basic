package main

import (
	"fmt"
)

func process(str string) bool {
	t := []rune(str)
	length := len(t)

	for i := range t {
		if i == length/2 {
			break
		}
		last := length - i - 1

		if t[i] != t[last] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Hello world!")
	var str string
	fmt.Scanf("%s", &str)
	fmt.Println(process(str))
}
