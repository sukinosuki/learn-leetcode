package main

import "fmt"

func check(nums []int) bool {

	flag := false

	if nums[0] < nums[len(nums)-1] {
		flag = true
	}
	for i := 1; i < len(nums); i++ {

		if nums[i] < nums[i-1] {
			if flag {
				return false
			}
			flag = true
		}
	}

	return true
}

func main() {

	// true
	//nums := []int{3, 4, 5, 1, 2}
	nums := []int{2, 1, 3, 4}

	res := check(nums)

	fmt.Printf("res: %v\n", res)

}
