package main

import "fmt"

// select the biggest number.
func ssort(a []int) {
	for i := 0; i < len(a); i++ {
		max := i
		for j := i + 1; j < len(a); j++ {
			if a[j] > a[max] {
				max = j
			}
		}
		a[i], a[max] = a[max], a[i]
	}
}

func main() {
	var a = []int{4, 2, 1, 65, 2132, 232, 34, 45, 15}
	ssort(a)
	fmt.Println(a)
}
