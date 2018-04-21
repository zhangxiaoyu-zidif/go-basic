package main

import "fmt"

func test(a, b int) int {
	result := func(a1 int, b1 int) int {
		return a1 + b1
	}(a, b)
	// 如果不加（a,b），这个reuslt就是一个函数
	return result
}

func main() {
	fmt.Println(test(10, 20))
}
