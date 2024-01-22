package main

import "fmt"

func sortArrayByParityII(nums []int) []int {

	l1 := 0
	l2 := 1

	ans := make([]int, len(nums))

	for _, num := range nums {

		if num%2 == 1 {
			ans[l2] = num
			l2 += 2
		} else {
			ans[l1] = num
			l1 += 2
		}
	}

	return ans
}

func main() {
	nums := []int{4, 2, 5, 7}
	res := sortArrayByParityII(nums)

	fmt.Printf("res: %v\n", res)
}
