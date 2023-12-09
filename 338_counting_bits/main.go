package main

import (
	"fmt"
)

func countBits(n int) []int {

	ans := make([]int, n+1)

	for i := 0; i <= n; i++ {
		num := i
		count := 0
		for num > 0 {
			if num&1 == 1 {
				count++
			}
			num >>= 1
		}

		ans[i] = count
	}

	return ans
}

func main() {
	n := 5
	res := countBits(n)

	fmt.Printf("res: %v\n", res)
}
