package main

import "fmt"

func sortArrayByParityII(nums []int) []int {

	for i, j := 0, 1; i < len(nums); i += 2 {

		// 如果偶数位值为奇数，查找奇数位值为偶数的值
		if nums[i]%2 == 1 {
			// 在查找奇数位值为偶数的过程中，如果奇数位值还是奇数，循环这个过程
			for nums[j]%2 == 1 {
				j += 2
			}

			// 偶数位的奇数与奇数位的偶数换值
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return nums
}

func main() {
	//nums := []int{5, 2, 4, 7}
	nums := []int{1, 2, 3, 3, 2, 3, 0, 4}
	res := sortArrayByParityII(nums)

	fmt.Printf("res: %v\n", res)
}
