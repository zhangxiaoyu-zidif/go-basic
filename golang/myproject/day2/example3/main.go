package main

import (
	"fmt"
	"time"
)

const (
	man    = 1
	female = 2
)

func main() {
	for {
		secondCount := time.Now().Unix()
		if secondCount%female == 0 {
			fmt.Println("female")
		} else {
			fmt.Println("male")
		}
		time.Sleep(time.Second)
	}
}
