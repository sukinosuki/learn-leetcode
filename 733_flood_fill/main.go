package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	arr := [][]int{
		{sr, sc},
	}

	prevColor := image[sr][sc]
	if prevColor == color {
		return image
	}

	image[sr][sc] = color

	m := len(image)
	n := len(image[0])
	for len(arr) > 0 {

		cells := arr[:]
		arr = nil

		for _, cell := range cells {
			x := cell[0]
			y := cell[1]

			// 上
			if x > 0 && image[x-1][y] == prevColor {
				image[x-1][y] = color
				arr = append(arr, []int{x - 1, y})
			}
			// 下
			if x < m-1 && image[x+1][y] == prevColor {
				image[x+1][y] = color
				arr = append(arr, []int{x + 1, y})
			}
			// 左
			if y > 0 && image[x][y-1] == prevColor {
				image[x][y-1] = color
				arr = append(arr, []int{x, y - 1})
			}
			// 右
			if y < n-1 && image[x][y+1] == prevColor {
				image[x][y+1] = color
				arr = append(arr, []int{x, y + 1})
			}
		}
	}

	return image
}

func main() {
	//image := [][]int{
	//	{1, 1, 1},
	//	{1, 1, 0},
	//	{1, 0, 1},
	//}
	//
	//res := floodFill(image, 1, 1, 2)
	image := [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}

	res := floodFill(image, 0, 0, 2)
	fmt.Printf("res: %v\n", res)
}
