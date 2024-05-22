package main

import "fmt"

func findLucky(arr []int) int {

	cnt := make(map[int]int)
	for _, num := range arr {
		cnt[num]++
	}
	maxNum := -1
	for k, v := range cnt {
		if k == v && k > maxNum {
			maxNum = k
		}
	}

	return maxNum
}

func main() {

	// 2
	//arr := []int{2, 2, 3, 4}
	arr := []int{1, 2, 2, 3, 3, 3}
	res := findLucky(arr)

	fmt.Printf("res: %v\n", res)
}
