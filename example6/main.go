package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
		return
	}

	strSlice := strings.Split(string(result), "+")

	numA := strings.TrimSpace(strSlice[0])
	numB := strings.TrimSpace(strSlice[1])

	fmt.Println(add(numA, numB))
}

func add(numA, numB string) (result string) {
	lenA := len(numA)
	lenB := len(numB)
	left := 0

	if lenA == 0 && lenB == 0 {
		result = "0"
	}
	if lenA < lenB {
		numA, numB = func(str1, str2 string) (string, string) {
			return str2, str1
		}(numA, numB)
	}
	lenA = len(numA)
	lenB = len(numB)

	for lenB > 0 {
		c1 := numA[lenA-1] - '0'
		c2 := numB[lenB-1] - '0'

		sum := int(c1) + int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := sum%10 + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		lenA--
		lenB--
	}
	for lenA > 0 {
		c1 := numA[lenA] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := sum%10 + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		lenA--
	}
	if left == 1 {
		result = fmt.Sprintf("1%s", result)
	}
	return
}
