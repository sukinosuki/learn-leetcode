package main

import (
	"fmt"
)

func countGoodTriplets(arr []int, a int, b int, c int) int {

	ans := 0
	n := len(arr)
	for i := 0; i < n; i++ {

		for j := i + 1; j < n; j++ {
			if abs(arr[i]-arr[j]) <= a {
				for k := j + 1; k < n; k++ {
					if abs(arr[i]-arr[k]) <= c && abs(arr[j]-arr[k]) <= b {
						ans++
					}
				}
			}
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
	arr := []int{3, 0, 1, 1, 9, 7}
	a := 7
	b := 2
	c := 3
	res := countGoodTriplets(arr, a, b, c)
	fmt.Printf("res: %v\n", res)
}
