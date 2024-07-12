package main

import "fmt"

func countQuadruplets(nums []int) int {

	ans := 0

	n := len(nums)
	for i := 0; i < n; i++ {

		for j := i + 1; j < n; j++ {

			for z := j + 1; z < n; z++ {

				for x := z + 1; x < n; x++ {

					if nums[i]+nums[j]+nums[z] == nums[x] {
						ans++
					}
				}
			}
		}
	}

	return ans
}

func main() {
	//1
	//nums := []int{1, 2, 3, 6}
	nums := []int{1, 1, 1, 3, 5}
	res := countQuadruplets(nums)

	fmt.Printf("res: %v\n", res)
}
