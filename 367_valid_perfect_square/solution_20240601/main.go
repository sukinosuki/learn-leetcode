package main

import "fmt"

func isPerfectSquare(num int) bool {

	if num == 1 {
		return true
	}
	l := 0
	r := num / 2
	for l <= r {

		mid := l + (r-l)/2
		if mid*mid == num {
			return true
		}

		if mid*mid > num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}

func main() {

	num := 9
	res := isPerfectSquare(num)

	fmt.Printf("res: %v\n", res)
}
