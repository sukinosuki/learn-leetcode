package main

import "fmt"

func kLengthApart(nums []int, k int) bool {
	n := len(nums)

	l := 0

	for l < n && nums[l] == 0 {
		l++
	}

	l2 := l + 1
	for l2 < n {
		if nums[l2] == 0 {
			l2++
			continue
		}
		if l2-l-1 < k {
			return false
		}

		l = l2
		l2++
	}

	return true
}

func main() {

	// true
	//nums := []int{1, 0, 0, 0, 1, 0, 0, 1}
	//k:=2
	// false
	//nums := []int{1, 0, 0, 1, 0, 1}
	//k := 2

	// true
	//nums := []int{1, 1, 1, 1, 1}
	//k := 0

	// true
	//nums := []int{0, 1, 0, 1}
	//k := 1

	nums := []int{0, 0, 0}
	k := 2
	res := kLengthApart(nums, k)

	fmt.Printf("res: %v\n", res)
}
