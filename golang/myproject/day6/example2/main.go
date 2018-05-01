package main

import (
	"fmt"
)

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("Read data")
}

func (f *File) Write() {
	fmt.Println("Write data")
}

func Test(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main() {
	var f File
	Test(&f)

	var b interface{}
	var x *File
	b = x
	v, ok := b.(ReadWriter)
	fmt.Println(v, ok)
	v.Read()
	v.Write()
}
