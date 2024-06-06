package main

import "fmt"

func findKthPositive(arr []int, k int) int {

	index := 0
	num := 1
	n := len(arr)

	i := 0
	for index < k {

		for num < arr[i] {
			index++

			if index == k {
				return num
			}

			num++
		}

		i++

		if i == n {
			for index < k {
				num++
				index++
			}
			return num
		}
		num++
	}

	return num
}

func main() {
	//arr := []int{2, 3, 4, 7, 11}
	//k := 5
	arr := []int{1, 2, 3, 4}
	k := 2

	res := findKthPositive(arr, k)

	fmt.Printf("res: %v\n", res)
}
