package main

import "fmt"

func tribonacci(n int) int {

	var helper func(n int) int
	m := make(map[int]int)

	helper = func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 || n == 2 {
			return 1
		}

		n1Res, ok := m[n-1]
		if !ok {
			n1Res = helper(n - 1)
			m[n-1] = n1Res
		}
		n2Res, ok := m[n-2]
		if !ok {
			n2Res = helper(n - 2)
			m[n-2] = n2Res
		}
		n3Res, ok := m[n-3]
		if !ok {
			n3Res = helper(n - 3)
			m[n-3] = n3Res
		}

		return n1Res + n2Res + n3Res
	}

	res := helper(n)

	return res
}

func main() {
	//n := 3 // 2
	//n := 4 // 4
	n := 5 // 7
	//n := 6 // 13
	//n := 25 // 4

	res := tribonacci(n)

	fmt.Printf("res: %v\n", res)
}
