package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {

	ans := []string{}

	if len(nums) == 0 {
		return ans
	}
	if len(nums) == 1 {
		ans = append(ans, strconv.Itoa(nums[0]))
		return ans
	}

	i := 0
	j := 1
	l := 0

	for l < len(nums) {
		if j >= len(nums) || nums[i] != nums[j]-1 {
			if i > l {
				ans = append(ans, fmt.Sprintf("%d->%d", nums[l], nums[i]))
			} else {
				ans = append(ans, strconv.Itoa(nums[i]))
			}

			l = j
			i = j
			j = i + 1
		} else {
			j++
			i++
		}
	}

	return ans
}

func main() {

	nums := []int{0, 1, 2, 4, 5, 7}
	res := summaryRanges(nums)

	fmt.Printf("res: %v\n", res)
}
