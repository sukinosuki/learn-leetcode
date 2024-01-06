package main

import (
	"fmt"
)

func dominantIndex(nums []int) int {

	subMaxIndex := -1
	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			subMaxIndex = maxIndex
			maxIndex = i
			continue
		}
		if subMaxIndex == -1 || nums[i] > nums[subMaxIndex] {
			subMaxIndex = i
		}
	}

	if subMaxIndex != -1 && nums[maxIndex] >= nums[subMaxIndex]*2 {
		return maxIndex
	}

	return -1

}

func main() {
	//nums := []int{3, 6, 1, 0}
	nums := []int{0, 0, 3, 2}
	res := dominantIndex(nums)

	fmt.Printf("res: %v\n", res)
}
