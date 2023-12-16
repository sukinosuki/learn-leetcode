package main

import "fmt"

func arrangeCoins(n int) int {

	l := 0
	r := n
	ans := 0
	for l <= r {
		mid := l + (r-l)/2

		if (mid*(mid+1))/2 <= n {
			l = mid + 1
			ans = mid
		} else {
			r = mid - 1
		}

	}

	return ans
}

func main() {
	n := 5
	//n := 8

	stagger := arrangeCoins(n)

	fmt.Printf("stagger: %v\n", stagger)
}
