package main

func largestTriangleArea(points [][]int) float64 {
	maxX, maxY := 0, 0

	for _, arr := range points {
		maxX = max(maxX, arr[0])
		maxY = max(maxY, arr[1])
	}

	return (float64(maxX) * float64(maxY)) / 2
}

func max(n1, n2 int) int {

	if n1 > n2 {
		return n1
	}

	return n2

}

func main() {

}
