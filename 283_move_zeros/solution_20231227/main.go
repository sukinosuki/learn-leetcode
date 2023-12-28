package main

func moveZeroes(nums []int) {
	l1 := 0
	l2 := 0

	for l2 < len(nums) {

		if nums[l2] == 0 {
			l2++
			continue
		}

		nums[l1], nums[l2] = nums[l2], nums[l1]

		l1++
		l2++
	}

}

func main() {

}
