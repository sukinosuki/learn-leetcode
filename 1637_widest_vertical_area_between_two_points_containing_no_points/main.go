package main

import (
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {

		return points[i][0] < points[j][0]
	})

	ans := 0

	for i := 1; i < len(points); i++ {

		sub := abs(points[i][0] - points[i-1][0])
		if sub > ans {
			ans = sub
		}
	}

	return ans
}

func abs(n int) int {

	if n > 0 {
		return n
	}

	return -n
}
func main() {

}
