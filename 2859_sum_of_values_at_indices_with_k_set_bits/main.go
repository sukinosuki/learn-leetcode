package main

import "fmt"

func sumIndicesWithKSetBits(nums []int, k int) int {

	sum := 0
	for i := 0; i < len(nums); i++ {
		n := getCount1FromBitNumber(i)
		if k == n {
			sum += nums[i]
		}
	}
	return sum
}

func getCount1FromBitNumber(num int) int {
	count := 0
	for num > 0 {
		count++
		num &= num - 1
	}

	return count
}

func main() {
	nums := []int{5, 10, 1, 5, 2}
	res := sumIndicesWithKSetBits(nums, 1)

	fmt.Printf("res: %v\n", res)
}
