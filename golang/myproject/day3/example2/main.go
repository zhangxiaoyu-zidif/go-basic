package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006-1-02 15:04"))

	calcDuration()
}

func calcDuration() {
	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()
	fmt.Printf("cost: %d us\n", (end-start)/1000)
}

func test() {
	time.Sleep(time.Millisecond * 100)
}
