package main

import "fmt"

func sumOfTheDigitsOfHarshadNumber(x int) int {

	num := x
	sum := 0
	for num > 0 {

		sum += num % 10
		num /= 10
	}

	if x%sum == 0 {
		return sum
	}

	return -1
}

func main() {

	// 9
	//x := 18
	x := 23

	res := sumOfTheDigitsOfHarshadNumber(x)

	fmt.Printf("res: %v\n", res)
}
