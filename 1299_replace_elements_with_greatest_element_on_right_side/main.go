package main

import "fmt"

func replaceElements(arr []int) []int {

	n := len(arr)
	ans := make([]int, n)

	maxNum := -1

	for i := n - 1; i >= 0; i-- {
		num := arr[i]

		ans[i] = maxNum

		if num > maxNum {
			maxNum = num
		}
	}

	return ans
}

func main() {

	//arr := []int{17, 18, 5, 4, 6, 1}
	arr := []int{400}
	res := replaceElements(arr)
	fmt.Printf("res: %v\n", res)
}
