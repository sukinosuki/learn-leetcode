package main

import "fmt"

func binaryGap(n int) int {
	ans := 0

	prev := -1
	for n > 0 {
		if n&1 == 0 {
			if prev != -1 {
				prev++
			}
		}

		if n&1 == 1 {
			ans = max(ans, prev+1)
			prev = 0
		}

		n >>= 1
	}

	return ans
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2

}

func main() {
	//n := 22 // 2
	n := 8
	res := binaryGap(n)

	fmt.Printf("res: %v\n", res)
}
