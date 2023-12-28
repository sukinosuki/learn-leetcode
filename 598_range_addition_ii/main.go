package main

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}

	minX := ops[0][0]
	minY := ops[0][1]
	for i := 1; i < len(ops); i++ {
		if minX > ops[i][0] {
			minX = ops[i][0]
		}
		if minY > ops[i][1] {
			minY = ops[i][1]
		}
	}

	return minX * minY
}

func main() {

}
