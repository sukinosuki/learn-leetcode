package main

import "fmt"

func findGCD(nums []int) int {

	minNum := nums[0]
	maxNum := minNum

	for i := range nums {
		if i > 0 {

			if nums[i] < minNum {
				minNum = nums[i]
			}
			if nums[i] > maxNum {
				maxNum = nums[i]
			}
		}
	}

	return findGreatCommonDivisor(maxNum, minNum)
}

func findGreatCommonDivisor(n1, n2 int) int {

	if n2 == 0 {
		return n1
	}
	return findGreatCommonDivisor(n2, n1%n2)
	//for n2 != 0 {
	//	n1, n2 = n2, n1%n2
	//}
	//
	//return n1
}

func main() {
	nums := []int{2, 5, 6, 9, 10}
	res := findGCD(nums)

	fmt.Printf("res: %v\n", res)
}
