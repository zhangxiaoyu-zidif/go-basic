package main

import (
	"fmt"
)

type testdata struct {
	a int
	b string
	c bool
}

func main() {
	var a = 10
	str := fmt.Sprintf("a = %d", a)

	testdataExp := testdata{1, "abc", true}

	var point *int
	point = &a

	fmt.Printf("%T\n", str)
	fmt.Printf("%q\n", str)

	fmt.Printf("%T\n", testdataExp)
	fmt.Printf("%v\n", testdataExp)

	dataType := fmt.Sprintf("%T", testdataExp)
	fmt.Println(dataType)

	dataType = fmt.Sprintf("%T", point)

	fmt.Println(dataType)

	str2 := "hello world"
	subStr := str2[0:5]
	fmt.Println(subStr)

	substr := str2[6:]
	fmt.Println(substr)

	fmt.Println(converseString(str2))
	fmt.Println(converseStringByByte(str2))
}

func converseString(input string) string {
	var output string
	length := len(input)
	if length == 0 {
		return ""
	}
	for i := 0; i < length; i++ {
		//output += input[length-i-1 : length-i]
		output += fmt.Sprintf("%c", input[length-i-1])
	}
	return output
}

func converseStringByByte(input string) string {
	tmp := []byte(input)
	var result []byte
	length := len(tmp)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1])
	}
	return string(result)
}
