package main

import "fmt"

func mySqrt(x int) int {
	l := 0
	r := x
	ans := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}

func main() {
	x := 4

	res := mySqrt(x)
	fmt.Printf("res: %d\n", res)
}
