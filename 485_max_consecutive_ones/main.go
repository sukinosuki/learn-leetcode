package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {

	max := 0
	prevMax := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			prevMax++
		} else {
			if max < prevMax {
				max = prevMax
			}

			prevMax = 0
		}
	}
	if max < prevMax {
		max = prevMax
	}

	return max
}

func main() {
	//nums := []int{1, 1, 0, 1, 1, 1}
	nums := []int{1, 0, 1, 1, 0, 1}
	max := findMaxConsecutiveOnes(nums)

	fmt.Printf("max: %v\n", max)
}
