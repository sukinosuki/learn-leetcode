package main

import "fmt"

func canBeIncreasing(nums []int) bool {

	top := nums[0]
	cnt := 0
	l := 1
	// 	nums := []int{105, 924, 32, 968}
	for l < len(nums) {

		// 当前元素非递增
		if nums[l] <= top {
			cnt++
			// 元素1与元素0相比为非递增，top更新为元素1
			if l == 1 || nums[l] > nums[l-2] {
				top = nums[l]
			} else {
				top = nums[l-1]
			}
		} else {
			// top更新为当前元素
			top = nums[l]
		}

		l++
		if cnt > 1 {
			return false
		}
	}

	return true
}

func main() {
	// true
	//nums := []int{1, 2, 10, 5, 7}
	// true
	//nums := []int{100, 21, 100}
	// false
	//nums := []int{541, 783, 433, 744}
	// false
	//nums := []int{1, 1, 1}

	// true
	//nums := []int{1, 2, 3}
	// true
	//nums := []int{105, 924, 32, 968}

	// false
	nums := []int{2, 3, 1, 2}
	res := canBeIncreasing(nums)

	fmt.Printf("res: %v\n", res)
}
