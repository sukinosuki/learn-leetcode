package main

import "fmt"

func arithmeticTriplets(nums []int, diff int) int {

	i, j := 0, 1
	ans := 0
	for _, x := range nums[2:] {

		for nums[j]+diff < x {
			j++
		}
		if nums[j]+diff > x {
			continue
		}

		for nums[i]+diff*2 < x {
			i++
		}
		if nums[i]+diff*2 == x {
			ans++
		}
	}

	return ans
}

func main() {
	// 2
	//nums := []int{0, 1, 4, 6, 7, 10}
	//diff := 3
	nums := []int{4, 5, 6, 7, 8, 9}
	diff := 2
	res := arithmeticTriplets(nums, diff)

	fmt.Printf("res: %v\n", res)
}
