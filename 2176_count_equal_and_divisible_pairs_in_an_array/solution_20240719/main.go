package main

import "fmt"

func countPairs(nums []int, k int) int {

	ans := 0
	cnt := make(map[int][]int)
	for i, num := range nums {
		if arr, ok := cnt[num]; ok {

			for _, j := range arr {
				if i*j%k == 0 {
					ans++
				}
			}
			cnt[num] = append(cnt[num], i)
		} else {
			cnt[num] = []int{i}
		}
	}

	return ans
}

func main() {
	// 4
	//nums := []int{3, 1, 2, 2, 2, 1, 3}
	//k := 2
	//18
	nums := []int{5, 5, 9, 2, 5, 5, 9, 2, 2, 5, 5, 6, 2, 2, 5, 2, 5, 4, 3}
	k := 7
	res := countPairs(nums, k)

	fmt.Printf("res: %v\n", res)
}
