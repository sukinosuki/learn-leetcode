package main

import "fmt"

func sumZero(n int) []int {

	ans := make([]int, n)
	i := 0
	if n%2 == 1 {
		i = 1
		ans[0] = 0
	}

	for i < n {

		ans[i] = n - i
		ans[i+1] = -(n - i)
		i += 2
	}

	return ans
}

func main() {

	n := 8
	res := sumZero(n)
	fmt.Printf("res: %v\n", res)
}
