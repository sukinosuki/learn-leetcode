package main

import (
	"fmt"
)

func largestSumAfterKNegations(nums []int, k int) int {

	numbers := make([]int, 201)
	for _, v := range nums {
		numbers[v+100]++
	}
	i := 0
	for k > 0 {
		for numbers[i] == 0 { // 找到nums中最小的数字
			i++
		}
		numbers[i]--     // 此数字个数-1
		numbers[200-i]++ //其相反个+1
		if i > 100 {     // 若原最小数索引>100, 则新的最小索引应为200-i. (索引即nums[]数组的下标)
			i = 200 - i
		}
		k--
	}

	sum := 0
	for j := i; j < len(numbers); j++ {
		sum += (j - 100) * numbers[j] //j-100是数字大小，numbers[j]是该数字出现次数
	}

	return sum
}

func main() {
	// 5
	//nums := []int{4, 2, 3}
	//k := 1

	// 6
	nums := []int{3, -1, 0, 2}
	k := 3

	// 13
	//nums := []int{2, -3, -1, 5, -4}
	//k := 2

	// 22
	//nums := []int{-8, 3, -5, -3, -5, -2}
	//k := 6

	// 20
	//nums := []int{5, 6, 9, -3, 3}
	//k := 2
	res := largestSumAfterKNegations(nums, k)

	fmt.Printf("res: %d\n", res)
}
