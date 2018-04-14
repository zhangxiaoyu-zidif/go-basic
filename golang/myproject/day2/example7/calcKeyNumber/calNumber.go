package calckeynumber

func CalcPrimeNumber(from int, to int) int {
	var count = 0
	for j := from; j <= to; j++ {
		for i := 2; i <= j; i++ {
			if j%i == 0 && i != j {
				break
			}
			if i == j {
				count++
			}
		}
	}
	return count
}

func CalcNarcissisticNumber(from, to int) int {
	count := 0
	if from < 100 && to > 999 {
		return 0
	}
	var oneDigit, twoDigit, threeDigit int
	for i := from; i <= to; i++ {
		oneDigit = i % 10
		threeDigit = i / 100
		twoDigit = (i - threeDigit*100) / 10
		if calcTrible(oneDigit)+calcTrible(twoDigit)+calcTrible(threeDigit) == i {
			count++
		}
	}
	return count
}

func calcTrible(number int) int {
	return number * number * number
}

func calcFactorial(number int) int {
	var value = 1
	for i := 1; i <= number; i++ {
		value *= i
	}
	return value
}

func CalcFactorialSum(number int) int {
	sum := 0
	for i := 1; i <= number; i++ {
		sum += calcFactorial(i)
	}
	return sum
}
