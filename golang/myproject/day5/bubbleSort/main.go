package main

import (
	"fmt"
)

func bsort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}

func main() {
	var a = []int{4, 2, 1, 65, 2132, 232, 34, 45, 15}
	bsort(a)
	fmt.Println(a)
}
