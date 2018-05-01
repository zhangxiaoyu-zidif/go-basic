package main

import (
	"fmt"
	"time"
)

type Car struct {
	Name string
	Age  int
}

type Train struct {
	Car
	createTime time.Time
	int
}

type Carer interface {
	GetName() string
	Run()
	DiDi()
}

type BMW struct {
	Name string
}

func (p *BMW) GetName() string {
	return p.Name
}

func (p *BMW) Run() {
	fmt.Println(p.Name + " runs.")
}

func (p *BMW) DiDi() {
	fmt.Println(p.Name + " is a didi.")
}

func (train *Train) Run() {
	train.Name = "Train " + train.Name
	fmt.Println(train.Name + " is running!")
}

type BYD struct {
	Name string
}

func (p *BYD) GetName() string {
	return p.Name
}

func (p *BYD) Run() {
	fmt.Println(p.Name + " runs.")
}

func (p *BYD) DiDi() {
	fmt.Println(p.Name + " is a didi.")
}

type Test interface {
	Hello()
}

func (p *BMW) Hello() {
	fmt.Println(p.Name + " says Hello")
}

func main() {
	var train Train
	train.int = 300
	train.Car.Name = "test"
	fmt.Println(train)
	train.Run()
	fmt.Println(train.Name)

	// ---------------------
	var a interface{}
	fmt.Printf("%T\n", a)
	var b int
	a = b
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	//-------------------------
	var car Carer
	var bmw BMW
	bmw.Name = "BMW"
	car = &bmw
	car.Run()

	byd := &BYD{
		Name: "byd",
	}

	car = byd
	car.Run()

	var test Test
	test = &bmw
	test.Hello()
}
