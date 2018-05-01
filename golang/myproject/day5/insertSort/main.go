package main

import "fmt"

func isort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				break
			}
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func main() {
	var a = []int{4, 2, 1, 65, 2132, 232, 34, 45, 15}
	isort(a)
	fmt.Println(a)
}
