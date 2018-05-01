package main

import (
	"fmt"
)

func main() {
	var a interface{}
	var b int
	a = b
	c, ok := a.(int)
	if ok {
		fmt.Println("convert successfully.")
	}
	fmt.Printf("%T\n", c)

	var d int
	Test(d)

	classifier(d)
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%v is a float64\n", i)
		default:
			fmt.Printf("param #%d's type is unknown\n", i)
		}
	}
}

func Test(a interface{}) {
	fmt.Printf("%T\n", a)
}
