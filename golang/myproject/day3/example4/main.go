package main

import "fmt"

func main() {
	var a int = 0

	switch a {
	case -1:
		fmt.Println("a equals -1.")
	case 0:
		fmt.Println("a equals 0.")
		fallthrough
	case 10:
		fmt.Println("a equals 10.")
	default:
		fmt.Println("a equals default.")
	}
}
