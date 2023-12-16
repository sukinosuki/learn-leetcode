package main

func moveZeroes(nums []int) {
	l1 := 0
	l2 := 0
	for l2 < len(nums) {

		if nums[l2] != 0 {
			nums[l1] = nums[l2]
			l1++
		}
		l2++
	}

	for l2 < len(nums) {
		nums[l2] = 0
		l2++
	}

}

func main() {

}
