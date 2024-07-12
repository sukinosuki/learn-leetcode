package main

import "fmt"

func countKDifference(nums []int, k int) int {
	//m := make(map[int]int)
	//
	n := len(nums)
	ans := 0
	//for b := n - 1; b > 0; b-- {
	//	m[k+nums[b]] = 1
	//	m[abs(k-nums[b])] = 1
	//
	//	for i := 0; i < b; i++ {
	//		ans += m[nums[i]]
	//	}
	//
	//	delete(m, k+nums[b])
	//	delete(m, abs(k-nums[b]))
	//}
	//
	//return ans

	for j := n - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if nums[i]-nums[j] == k || nums[i]-nums[j] == -k {
				ans++
			}
		}
	}
	return ans
}

func abs(n int) int {
	if n > 0 {
		return n
	}

	return -n
}

func main() {
	//4
	//nums := []int{1, 2, 2, 1}
	//k := 1
	// 3
	nums := []int{3, 2, 1, 5, 4}
	k := 2
	res := countKDifference(nums, k)
	fmt.Printf("res: %v\n", res)
}
