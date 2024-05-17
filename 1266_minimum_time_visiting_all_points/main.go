package main

import "math"

func minTimeToVisitAllPoints(points [][]int) int {

	x0 := points[0][0]
	x1 := points[0][1]
	ans := 0
	for i := 0; i < len(points); i++ {
		y0, y1 := points[i][0], points[i][1]
		ans += max(int(math.Abs(float64(x0-y0))), int(math.Abs(float64(x1-y1))))

		x0 = y0
		x1 = y1
	}
	return ans
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}

func main() {

}
