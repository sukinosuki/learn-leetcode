package main

func runningSum(nums []int) []int {

	n := len(nums)
	ans := make([]int, n)

	ans[0] = nums[0]
	for i := 1; i < n; i++ {

		ans[i] = nums[i] + ans[i-1]
	}

	return ans
}

func main() {

}
