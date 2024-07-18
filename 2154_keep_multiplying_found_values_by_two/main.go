package main

import "fmt"

func findFinalValue(nums []int, original int) int {

	cnt := make(map[int]bool)
	for _, num := range nums {
		cnt[num] = true
	}

	for cnt[original] {
		original *= 2
	}

	return original
}

func main() {
	nums := []int{5, 3, 6, 1, 12}
	original := 3
	res := findFinalValue(nums, original)

	fmt.Printf("res: %v\n", res)
}
