package main

import (
	"fmt"
	"sort"
)

// func intersection(nums [][]int) []int {
//
//		cnt := make(map[int]int)
//
//		for index, arr := range nums {
//
//			for _, num := range arr {
//				cnt[num] |= 1 << index
//			}
//		}
//
//		var ans []int
//		mask := int(math.Pow(2, float64(len(nums)))) - 1
//		for k, v := range cnt {
//			if v == mask {
//				ans = append(ans, k)
//			}
//		}
//
//		sort.Ints(ans)
//
//		return ans
//	}
func intersection(nums [][]int) []int {

	cnt := make(map[int]int)

	n := len(nums)
	for index, arr := range nums {

		mask := 0
		for _, num := range arr {
			if mask&(1<<num) == 0 {
				mask = 1 << num
				cnt[num]++
			}
		}
		if index > 0 {
			tempCnt := cnt
			cnt = map[int]int{}
			for k, v := range tempCnt {
				if v == index+1 {
					cnt[k] = v
				}
			}
		}
	}

	var ans []int
	for k, v := range cnt {
		if v == n {
			ans = append(ans, k)
		}
	}

	sort.Ints(ans)

	return ans
}
func main() {
	nums := [][]int{
		{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6},
	}
	res := intersection(nums)

	fmt.Printf("res: %v\n", res)
}
