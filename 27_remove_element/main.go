package main

import "fmt"

func removeElement(nums []int, val int) int {

	n := len(nums)

	i := n - 1
	cnt := 0
	for i >= 0 && i >= cnt {
		if nums[i] != val {
			nums[cnt], nums[i] = nums[i], nums[cnt]
			cnt++
		} else {
			i--
		}
	}
	return cnt

}

func main() {

	//  [2 2 3 3]
	//nums := []int{3, 2, 2, 3}
	//value := 3

	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	value := 2

	res := removeElement(nums, value)

	fmt.Printf("res: %v\n", res)
}
