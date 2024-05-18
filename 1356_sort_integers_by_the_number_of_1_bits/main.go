package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {

	ans := make([]int, len(arr))
	list := make([]map[int]int, len(arr))

	for i := range arr {
		m := make(map[int]int)

		m[0] = count(arr[i])
		m[1] = arr[i]
		list[i] = m
	}

	sort.Slice(list, func(i, j int) bool {
		count1 := list[i][0]
		num1 := list[i][1]
		count2 := list[j][0]
		num2 := list[j][1]
		if count1 == count2 {

			return num1 < num2
		}

		return count1 < count2
	})

	for i := range list {
		ans[i] = list[i][1]
	}

	return ans
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
	//arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	arr := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	res := sortByBits(arr)

	fmt.Printf("res: %v\n", res)
}
