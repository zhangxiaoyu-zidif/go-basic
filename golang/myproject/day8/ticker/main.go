package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(2 * time.Second)
	for v := range t.C {
		fmt.Println("Hello", v)
	}
}
