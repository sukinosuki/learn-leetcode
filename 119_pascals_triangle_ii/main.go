package main

import "fmt"

func getRow(rowIndex int) []int {

	ans := make([]int, 1)

	for i := 0; i <= rowIndex; i++ {

		prev := ans[:]
		ans = make([]int, i+1)
		ans[0] = 1
		ans[i] = 1

		for j := 1; j < i; j++ {
			ans[j] = prev[j] + prev[j-1]
		}
	}

	return ans
}

func main() {
	ans := getRow(3)
	fmt.Printf("ans: %v\n", ans)
}
