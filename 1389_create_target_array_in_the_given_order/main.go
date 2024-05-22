package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {

	//var res = make([]int, len(nums))
	//
	//for k, i := range index {
	//	copy(res[i+1:], res[i:]) //Go 圣经里讲的，插入应该怎么插，就这么插，纠结不能用函数的，当然也可以循环移动
	//	res[i] = nums[k]         //插！
	//}
	//
	//return res
}

func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}

	res := createTargetArray(nums, index)

	fmt.Printf("res: %v\n", res)
}
