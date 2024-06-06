package main

import "fmt"

func mySqrt(x int) int {

	l := 0
	r := x

	for l <= r {

		mid := l + (r-l)/2
		if mid*mid == x {
			return mid
		}

		if mid*mid < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return r
}

func main() {

	x := 8
	res := mySqrt(x)

	fmt.Printf("res: %v\n", res)
}
