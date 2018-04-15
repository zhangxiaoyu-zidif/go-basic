package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("starting.")

	randomNum := rand.Intn(100)

	var inputNumber int
	fmt.Scanf("%d", &inputNumber)

	switch {
	case randomNum == inputNumber:
		fmt.Println("bingo.")
	case randomNum < inputNumber:
		fmt.Println("too big")
	case randomNum > inputNumber:
		fmt.Println("too small")
	}

}
