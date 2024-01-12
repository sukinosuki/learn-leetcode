package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {

	ans := make([][]int, len(image))
	n := len(image[0])
	for index := range ans {
		ans[index] = make([]int, n)
	}

	for index, arr := range image {

		for i := len(arr) - 1; i >= 0; i-- {

			if arr[i] == 0 {
				ans[index][n-i-1] = 1
			} else {
				ans[index][n-i-1] = 0

			}
		}

	}
	return ans
}

func main() {
	image := [][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}
	res := flipAndInvertImage(image)

	fmt.Printf("res: %v\n", res)
}
