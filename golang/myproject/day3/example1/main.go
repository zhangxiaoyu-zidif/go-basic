package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "https://www.baidu.com/"

	if checkValidUrl(str) {
		fmt.Println("This is a valid URL!")
	} else {
		fmt.Println("This is a invalid URL!")
	}

	execStings("oooabcLLLL abc0 kkkk")

	execStrconv()
}

func checkValidUrl(str string) bool {
	if strings.HasSuffix(str, "/") && strings.HasPrefix(str, "https://") {
		return true
	} else {
		return false
	}
}

func execStings(str string) {
	a := strings.Replace(str, "abc", "xxx", 1)
	fmt.Println(a)

	count := strings.Count(str, "abc")
	fmt.Println(count)

	a = strings.Trim("a123bc", "abc")
	fmt.Println(a)

	a = strings.Trim("a123bca", "a")
	fmt.Println(a)

	a = strings.TrimLeft("abc321", "abc")
	fmt.Println(a)

	a = strings.TrimRight("123abc", "abc")
	fmt.Println(a)

	slice := strings.Fields(str)
	fmt.Println(slice)

	slice = strings.Split(str, " ")
	fmt.Println(slice)

	stringSlice := []string{"abc", "123", "abc"}
	stringjoin := strings.Join(stringSlice, " ")
	fmt.Println(stringjoin)
}

func execStrconv() {
	a := 123
	str := strconv.Itoa(a)
	fmt.Println(str)
	fmt.Printf("%T\n", str)

	var st string
	fmt.Scanf("%s", &st)
	number, _ := strconv.Atoi(st)
	fmt.Println(number)
}
