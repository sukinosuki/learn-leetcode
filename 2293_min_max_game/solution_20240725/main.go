package main

import "fmt"

func minMaxGame(nums []int) int {

	arr := dfs(nums)

	return arr[0]

}

func dfs(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	temp := nums[:]
	nums = []int{}
	for i := 0; i < len(temp); i += 2 {
		if (i/2)%2 == 0 {

			nums = append(nums, min(temp[i], temp[i+1]))
		} else {
			nums = append(nums, max(temp[i], temp[i+1]))
		}
	}

	return dfs(nums)
}

func main() {
	// 1
	//nums := []int{1, 3, 5, 2, 4, 8, 2, 2}
	// 3
	//nums := []int{3}

	// 22
	nums := []int{70, 38, 21, 22}
	res := minMaxGame(nums)

	fmt.Printf("res: %v\n", res)
}
