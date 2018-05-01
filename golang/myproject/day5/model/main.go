package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	//S := new(Person)
	S := NewPerson("xiaoyu", 18)
	fmt.Println(S)
	fmt.Println(&S.Name)
}
