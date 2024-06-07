package main

func diagonalSum(mat [][]int) int {

	//rowNum := len(mat)
	//
	//l := 0
	//r := rowNum - 1
	//sum := 0
	//for l < r {
	//	// 00
	//	// 03
	//	// 30
	//	// 33
	//	sum += mat[l][l] + mat[l][r] + mat[r][l] + mat[r][r]
	//
	//	l++
	//	r--
	//}
	//if rowNum%2 == 1 {
	//	sum += mat[r][r]
	//}
	//
	//return sum

	rowNum := len(mat)

	l := 0
	r := rowNum - 1
	sum := 0

	// 结束条件为左边索引小于右边索引(即跳出循环时r一定等于l)
	for l < r {
		//      加左上角    加右上角     加左下角     加右下角
		sum += mat[l][l] + mat[l][r] + mat[r][l] + mat[r][r]

		// 往中间缩进
		l++
		r--
	}

	// 奇数的矩阵时，加上 l==r 中心位置的值
	if rowNum%2 == 1 {
		sum += mat[r][r]
	}

	return sum
}

func main() {

}
