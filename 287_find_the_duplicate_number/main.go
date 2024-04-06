package main

import "fmt"

func findDuplicate(nums []int) int {

	res := nums[0]

	for i := 1; i < len(nums); i++ {
		if i%2 == 1 {
			res -= nums[i]
		} else {
			res += nums[i]
		}
	}

	return res
}

func main() {
	//				0  1  2  3  4
	//nums := []int{1, 3, 4, 2, 2}//2
	// 1-3-2-4-2-4-2-4-2-4...
	//				0  1  2  3  4
	//nums := []int{3, 1, 3, 4, 2} //3
	// 3-4-2-3-4
	nums := []int{3, 1, 3, 4, 2} //3
	res := findDuplicate(nums)
	fmt.Println("res ", res)
}
