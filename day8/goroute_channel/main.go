package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println(i)
	}
}

func read(ch chan int) {
	for {
		var b int
		b = <-ch
		fmt.Println(b)
	}
}

func main() {
	intChan := make(chan int, 10)
	go write(intChan)
	go read(intChan)

	time.Sleep(10 * time.Second)
}
