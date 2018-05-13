package main

import (
	"fmt"
)

type Student struct {
	name string
}

// channel is FIFO
func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	intChan <- 1
	intChan <- 2
	intChan <- 3
	fmt.Println(len(intChan))
	fmt.Println(<-intChan)
	fmt.Println(len(intChan))

	//-----------------------
	var stuChan chan interface{}
	stuChan = make(chan interface{}, 10)
	stu := Student{name: "XXOO"}
	stuChan <- &stu

	var stu01 interface{}
	stu01 = <-stuChan
	fmt.Println(stu01)

	//var stu02 Student
	stu02, ok := stu01.(*Student)
	if ok {
		fmt.Println(stu02)
	}
}
