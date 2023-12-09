package main

import "fmt"

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)

	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}
	return NumArray{
		sums: sums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {

	return this.sums[right+1] - this.sums[left]
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	numArray := Constructor(nums)

	fmt.Printf("sums: %v\n", numArray)
}
