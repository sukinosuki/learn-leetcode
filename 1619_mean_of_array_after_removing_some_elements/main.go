package main

import (
	"fmt"
	"sort"
)

func trimMean(arr []int) float64 {

	n := int(float64(len(arr)) * 0.05)

	sort.Ints(arr)

	sum := 0
	for i := n; i < len(arr)-n; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr)-n*2)
}

func main() {

	// 2
	//arr := []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3}
	arr := []int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4}
	res := trimMean(arr)

	fmt.Printf("res: %v\n", res)
}
