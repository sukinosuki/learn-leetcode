package main

import "fmt"

func prefixesDivBy5(nums []int) []bool {

	n := 0
	res := make([]bool, len(nums))

	for j, v := range nums {

		n = (n<<1 + v) % 5
		res[j] = n == 0
	}

	return res
}

func main() {
	//nums := []int{0, 1, 1} // [true false false]
	//nums := []int{1, 1, 1} // [false false false]
	nums := []int{1, 0, 1} // [false false true]
	//nums := []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0} // [false false false]
	res := prefixesDivBy5(nums)

	fmt.Printf("res: %v\n", res)
}
