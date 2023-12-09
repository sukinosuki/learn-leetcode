package main

import "fmt"

func addDigits(num int) int {

	sum := 0
	copyNum := num
	for copyNum > 0 {

		for copyNum > 0 {
			delivery := copyNum % 10
			sum += delivery
			copyNum /= 10
		}
		if sum >= 10 {
			copyNum = sum
			sum = 0
		}
	}

	return sum
}

func main() {
	num := 38
	sum := addDigits(num)
	fmt.Printf("sum: %v\n", sum)
}
