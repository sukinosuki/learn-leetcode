package main

import "fmt"

func divideArray(nums []int) bool {

	cnt := make(map[int]int)

	for _, num := range nums {
		cnt[num]++
	}
	for _, v := range cnt {
		if v%2 == 1 {
			return false
		}
	}

	return true
}

func main() {
	nums := []int{3, 2, 3, 2, 2, 2}
	res := divideArray(nums)

	fmt.Printf("res: %v\n", res)
}
