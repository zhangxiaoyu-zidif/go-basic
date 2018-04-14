package main

import (
	"fmt"
	a "myproject/day2/example2/add"
)

const (
	abc = "abc"
)

func init() {
	fmt.Println("initialized the function.")
}

func main() {
	fmt.Println("go run main.")
	fmt.Println(a.Age, a.Name, abc)
}
