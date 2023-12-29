package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	max := math.MinInt
	k2 := k - 1
	for k2 < len(nums) {
		sum := nums[k2]

		if k2 >= k {
			sum -= nums[k2-k]
		}
		if max < sum {
			max = sum
		}
		k2++
	}

	return float64(max) / float64(k)
}

func main() {
	//nums := []int{1, 12, -5, -6, 50, 3}
	nums := []int{0, 1, 1, 3, 3}
	maxAvg := findMaxAverage(nums, 4)

	fmt.Printf("maxAvg: %v\n", maxAvg)
}
