package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {

	sort.Slice(arr, func(i, j int) bool {

		countI := count(arr[i])
		countJ := count(arr[j])
		if countI == countJ {
			return arr[i] < arr[j]
		}

		return countI < countJ
	})

	return arr
}

func count(n int) int {
	ans := 0
	for n > 0 {
		ans++
		n = n & (n - 1)
	}

	return ans
}

func main() {
	//  [0 1 2 4 8 3 5 6 7]
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	//arr := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	res := sortByBits(arr)

	fmt.Printf("res: %v\n", res)
}
