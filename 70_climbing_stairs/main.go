package main

import "fmt"

// 滚动数组
func climbStairs(n int) int {

	// 1	2	3	4	5
	// 1	2	3	5	8
	// fn(n) = fn(n-1) + fn(n-2)
	if n == 1 {
		return 1
	}
	a, b := 1, 2

	for i := 3; i <= n; i++ {
		sum := a + b
		a = b
		b = sum
	}

	return b
}

func main() {
	res := climbStairs(4)

	fmt.Printf("res: %d\n", res)
}
