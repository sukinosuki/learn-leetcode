package main

import "sort"

func findDisappearedNumbers(nums []int) []int {
	ans := make([]int, 0)
	sort.Ints(nums)

	for index, value := range nums {
		if value-index-1 != 0 {
			ans = append(ans, index+1)
		}
	}

	return ans
}

func main() {

}
