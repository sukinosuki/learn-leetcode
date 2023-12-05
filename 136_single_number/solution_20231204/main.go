package main

import "fmt"

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		fmt.Printf("num: %b\n", num)
		single ^= num
		fmt.Printf("single: %b\n", single)
	}

	return single
}

func main() {

	nums := []int{4, 1, 2, 1, 2}
	num := singleNumber(nums)

	fmt.Printf("num: %v\n", num)
}
