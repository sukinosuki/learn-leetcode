package main

import (
	"fmt"
)

func isBoomerang(points [][]int) bool {

	// 假设3个点坐标(x1,y1),(x2,y2),(x3,y3)，不在一条直线，只要斜率不一样即可。
	//(y2 - y1) / (x2 -x1) != (y3 - y1) / (x3 - x1)
	//除法存在精度问题，换算成乘法(y2 - y1) * (x3 -x1) != (y3 - y1) * (x2 - x1)

	//      (y2 - y1) * (x3 - x1)
	p1 := (points[1][1] - points[0][1]) * (points[2][0] - points[0][0])
	//      (x2 - x1) * (y3 - y1)
	p2 := (points[1][0] - points[0][0]) * (points[2][1] - points[0][1])

	return p1 != p2
}

func main() {
	// true
	//points := [][]int{
	//	{1, 1},
	//	{2, 2},
	//	{3, 3},
	//}

	// false
	//points := [][]int{
	//	{1, 1},
	//	{2, 3},
	//	{3, 2},
	//}

	// true
	points := [][]int{
		{0, 0},
		{0, 2},
		{2, 1},
	}
	res := isBoomerang(points)
	fmt.Printf("res: %v\n", res)
}
