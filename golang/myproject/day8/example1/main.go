package main

import (
	"fmt"
	"strings"
)

func main() {
	var a = "A"
	for i := 1; i < 6; i++ {
		str := strings.Repeat(a, i)
		str += "\n"
		fmt.Println(str)
	}

	str := "中国"
	fmt.Println(len([]byte(str)))
}
