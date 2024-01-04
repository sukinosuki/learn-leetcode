package main

import "fmt"

func pivotIndex(nums []int) int {

	n := len(nums)
	l := 0
	r := n - 1

	leftSum := 0
	rightSum := 0
	mid := l + (r-l)/2

	for i := 0; i < mid; i++ {
		leftSum += nums[i]
	}
	for i := mid + 1; i < n; i++ {
		rightSum += nums[i]
	}

	for leftSum != rightSum {
		if mid == 0 || mid == n-1 {
			return -1
		}

		if leftSum < rightSum {
			leftSum += nums[mid]
			rightSum -= nums[mid+1]
			mid++
			if leftSum > rightSum {
				return -1
			}
		} else {
			leftSum -= nums[mid-1]
			rightSum += nums[mid]
			mid--
			if leftSum < rightSum {
				return -1
			}
		}
	}

	return mid
}

func helper(leftSum, rightSum int, nums []int, mid int) int {
	n := len(nums)
	for leftSum != rightSum {
		if mid == 0 || mid == n-1 {
			return -1
		}

		leftSum += nums[mid]
		rightSum -= nums[mid+1]
		mid++
		if leftSum > rightSum {
			return -1
		}
	}

	return mid
}

func main() {
	//nums := []int{1, 7, 3, 6, 5, 6} // 3
	//nums := []int{1, 2, 3} // 3
	//nums := []int{2, 1, -1} // 3
	//nums := []int{-1, -1, -1, -1, -1, 1}
	nums := []int{-1, -1, -1, -1, 0, 1}
	index := pivotIndex(nums)
	fmt.Printf("index: %v\n", index)
}
