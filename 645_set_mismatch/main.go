package main

import "fmt"

func findErrorNums(nums []int) []int {
	arr := make([]int, len(nums))
	ans := make([]int, 2)
	for _, v := range nums {

		arr[v-1]++
	}

	for i, v := range arr {
		if v == 0 {
			ans[1] = i + 1
		}
		if v > 1 {
			ans[0] = i + 1
		}
	}

	return ans
}

func main() {
	nums := []int{1, 2, 2, 4}
	res := findErrorNums(nums)

	fmt.Printf("res: %v\n", res)
}
