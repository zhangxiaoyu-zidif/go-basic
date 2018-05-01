package main

import (
	"fmt"
)

func main() {
	var intLink Link
	for i := 0; i < 10; i++ {
		intLink.InsertHead(i)
	}

	intLink.Trans()

	fmt.Println()
	
	var intLink2 Link
	for i := 0; i < 10; i++ {
		intLink2.InsertTail(i)
	}

	intLink2.Trans()
}
