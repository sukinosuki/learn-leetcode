package main

import "fmt"

func plusOne(digits []int) []int {

	r := len(digits) - 1
	carry := 0
	for r >= 0 {
		sum := digits[r] + carry
		if r == len(digits)-1 {
			sum += 1
		}
		//carry = sum / 10
		digits[r] = sum % 10
		if sum > 9 {
			carry = 1
		} else {
			carry = 0
		}

		r--
	}
	if carry > 0 {
		res := []int{carry}
		res = append(res, digits...)

		return res
	}

	return digits
}

func main() {
	//digits := []int{1, 2, 3}
	digits := []int{9, 9, 9}

	res := plusOne(digits)

	fmt.Printf("res : %d\n", res)
}
