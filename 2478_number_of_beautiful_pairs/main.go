package main

import "fmt"

func countBeautifulPairs(nums []int) int {

	ans := 0

	for i := 0; i < len(nums); i++ {

		num1 := nums[i]
		for num1 > 9 {
			num1 /= 10
		}

		for j := i + 1; j < len(nums); j++ {

			num2 := nums[j] % 10

			k := max(num1, num2)

			for k > 1 && (num1%k != 0 || num2%k != 0) {
				k--
			}
			if k == 1 {
				ans++
			}

		}
	}

	return ans
}

func main() {

	// 5
	//nums := []int{2, 5, 4, 1}
	nums := []int{11, 21, 12}
	res := countBeautifulPairs(nums)

	fmt.Printf("res: %v\n", res)
}
