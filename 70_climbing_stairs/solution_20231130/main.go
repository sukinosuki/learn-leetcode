package main

import "fmt"

// 动态规划
func climbStairs(n int) int {

	if n == 1 {
		return 1
	}

	arr := make([]int, n)
	arr[0] = 1
	arr[1] = 2
	for i := 2; i < n; i++ {

		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n-1]
}

func main() {
	x := 1

	res := climbStairs(x)
	fmt.Printf("res: %d\n", res)
}
