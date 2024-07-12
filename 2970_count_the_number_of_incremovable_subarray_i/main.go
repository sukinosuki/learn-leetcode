package main

import "fmt"

func incremovableSubarrayCount(nums []int) int {

	ans := 0

	n := len(nums)

	if n == 2 {
		return 3
	}
	l := 0
	r := n - 1

	count := 1
	for r >= 0 {
		if r == 0 {
			ans += count
			break
		}

		if nums[r] == nums[r-1] {
			ans++
			r--
			continue
		}
		if nums[r] < nums[r-1] {
			if r == n-1 {
				ans++
			} else {
				ans += count
			}
			break
		}

		ans += count
		count++

		r--
	}

	if r == 0 {
		return ans
	}

	count = 1
	for l < r {
		if nums[l] >= nums[l+1] {
			if l == 0 {
				ans++
			} else {
				ans += count
			}
			break
		}

		ans += count
		count++
		l++
	}

	if l < r && nums[l] < nums[r] {
		size := l + 1 + n - 1 - r
		for size > 0 {
			ans += size
			size--
		}
	}

	return ans
}

func main() {
	// 10
	//nums := []int{
	//	1, 2, 3, 4,
	//}
	// 7
	//nums := []int{
	//	6, 5, 7, 8,
	//}
	// 12?
	//nums := []int{
	//	3, 4, 6, 5, 7, 8,
	//}
	// 3
	//nums := []int{
	//	8, 7, 6, 6,
	//}
	// 3
	//nums := []int{
	//	4, 2,
	//}

	// 3
	//nums := []int{
	//	1, 2,
	//}
	nums := []int{
		2, 1, 6,
	}
	res := incremovableSubarrayCount(nums)

	fmt.Printf("res: %v\n", res)
}
