package main

import "fmt"

func distinctDifferenceArray(nums []int) []int {

	m := make(map[int]int)
	m2 := make(map[int]int)
	n := len(nums)
	ans := make([]int, n)

	for _, v := range nums {

		m[v]++
	}

	for i, v := range nums {
		count := m[v]
		if count == 1 {
			delete(m, v)
		} else {
			m[v]--
		}

		m2[v]++
		ans[i] = len(m2) - len(m)
	}

	return ans
}

func main() {
	//nums := []int{1, 2, 3, 4, 5} //  [-3 -1 1 3 5]
	//            3  3  3  2  1
	//            1  2  2  3  3
	nums := []int{3, 2, 3, 4, 2} //  [-2 -1 0 2 3]
	res := distinctDifferenceArray(nums)
	fmt.Printf("res: %v\n", res)
}
