package main

import "fmt"

func maximumPrimeDifference(nums []int) int {

	arr := make([]int, 101)

	n := len(nums)
	l := 0
	r := n - 1

	for l < n {

		if arr[nums[l]] == 0 {
			if isPrime(nums[l]) {
				break
			} else {
				arr[nums[l]] = 1
			}
		}
		l++
	}
	if l == n {
		return 0
	}

	for r > l {
		if arr[nums[r]] == 0 {
			if isPrime(nums[r]) {
				break
			} else {
				arr[nums[r]] = 1
			}
		}
		r--
	}

	if r == l {
		return 0
	}

	return r - l
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	// 3
	//nums := []int{4, 2, 9, 5, 3}
	// 0
	//nums := []int{4, 8, 2, 8}
	nums := []int{1, 7}
	res := maximumPrimeDifference(nums)

	fmt.Printf("res: %v\n", res)
}
