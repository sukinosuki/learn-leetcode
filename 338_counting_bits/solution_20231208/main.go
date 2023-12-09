package main

import "fmt"

func countBits(n int) []int {

	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		count := 0
		for x := i; x > 0; x = x & (x - 1) {
			count++
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
