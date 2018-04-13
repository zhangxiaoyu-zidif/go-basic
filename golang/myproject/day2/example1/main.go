package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world, let's change the world!")
	filename := "C:\\abc.txt"
	createFile(filename)
	writeFile(filename)
	readfile(filename)
}