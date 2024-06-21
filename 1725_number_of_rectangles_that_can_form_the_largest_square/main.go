package main

import "fmt"

func countGoodRectangles(rectangles [][]int) int {

	ans := 0
	maxLen := 0

	for _, arr := range rectangles {

		minNum := min(arr[0], arr[1])
		if minNum == maxLen {
			ans++
			continue
		}

		if minNum > maxLen {
			ans = 1
			maxLen = minNum
		}
	}

	return ans
}

func main() {
	rectangles := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}
	res := countGoodRectangles(rectangles)

	fmt.Printf("res: %v\n", res)
}
