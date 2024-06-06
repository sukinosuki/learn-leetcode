package main

import "fmt"

func searchRange(nums []int, target int) []int {

	ans := []int{-1, -1}

	l := 0
	n := len(nums)
	r := n

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] == target {

			midL := mid
			for midL > 0 && nums[midL-1] == target {
				midL--
			}

			ans[0] = midL
			midR := mid
			for midR < n-1 && nums[midR+1] == target {
				midR++
			}
			ans[1] = midR

			return ans
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1

		}
	}

	return ans
}

func main() {
	// 3 4
	//nums := []int{5, 7, 7, 8, 8, 10}
	nums := []int{8, 8, 8, 8, 8, 8}
	target := 8

	res := searchRange(nums, target)

	fmt.Printf("res: %v\n ", res)
}
