package main

import (
	"fmt"
	"sort"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {

	sort.Ints(arr1)
	sort.Ints(arr2)

	i := len(arr1) - 1
	ans := 0
	for i >= 0 {

		num := arr1[i]
		ans++
		for j := len(arr2) - 1; j >= 0; j-- {

			sub := num - arr2[j]
			if sub < 0 {
				sub = -sub
			}

			if sub <= d {
				ans--
				break
			}
		}
		i--
	}

	return ans
}

func main() {
	// 2
	//d := 2
	//arr1 := []int{4, 5, 8}
	//arr2 := []int{10, 9, 1, 8}

	// 2
	//d := 3
	//arr1 := []int{1, 4, 2, 3}
	//arr2 := []int{-4, -3, 6, 10, 20, 30}

	// 1
	d := 6
	arr1 := []int{2, 1, 100, 3}
	arr2 := []int{-5, -2, 10, -3, 7}
	res := findTheDistanceValue(arr1, arr2, d)

	fmt.Printf("res: %v\n", res)
}
