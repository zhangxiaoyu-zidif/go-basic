package main

import (
	"fmt"
	"myproject/day2/example7/calckeynumber"
)

func main() {
	fmt.Println("calc the key numbers as below:")
	var a int
	a = calckeynumber.CalcPrimeNumber(100, 200)
	fmt.Println(a)

	a = calckeynumber.CalcNarcissisticNumber(100, 999)
	fmt.Println(a)

	a = calckeynumber.CalcFactorialSum(3)
	fmt.Println(a)
}
