package main

import (
	"fmt"
	"sort"
)

type KthLargest struct {
	nums         []int
	k            int
	largest3Nums []int
}

func Constructor(k int, nums []int) KthLargest {

	sort.Ints(nums)
	size := k
	arr := []int{}
	n := len(nums)

	for i, j := size-1, n-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		arr = append(arr, nums[j])
	}
	sort.Ints(arr)

	return KthLargest{
		k:            k,
		nums:         nums,
		largest3Nums: arr,
	}
}

func (this *KthLargest) Add(val int) int {
	this.nums = append(this.nums, val)
	if len(this.largest3Nums) < this.k || val > this.largest3Nums[0] {
		this.largest3Nums = append(this.largest3Nums, val)
		sort.Ints(this.largest3Nums)
		if len(this.largest3Nums) > this.k {
			this.largest3Nums = this.largest3Nums[1:]
		}
	}

	return this.largest3Nums[0]
}

func main() {
	//kth := Constructor(3, []int{4, 5, 8, 2})
	//
	//fmt.Println(kth.Add(3))
	//fmt.Println(kth.Add(5))
	//fmt.Println(kth.Add(10))
	//fmt.Println(kth.Add(9))
	//fmt.Println(kth.Add(4))

	//kth := Constructor(1, []int{})
	//fmt.Println(kth.Add(-3))
	//fmt.Println(kth.Add(-2))
	//fmt.Println(kth.Add(-4))
	//fmt.Println(kth.Add(0))
	//fmt.Println(kth.Add(4))
	kth := Constructor(2, []int{0})
	fmt.Println(kth.Add(-1))
	fmt.Println(kth.Add(1))
	fmt.Println(kth.Add(-2))
	fmt.Println(kth.Add(-4))
	fmt.Println(kth.Add(3))

}
