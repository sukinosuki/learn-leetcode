package main

import "sort"

func kWeakestRows(mat [][]int, k int) []int {

	arr := make([]float64, len(mat))

	m := make(map[float64]int)

	for i := range mat {
		sum := 0
		for _, num := range mat[i] {
			sum += num
		}
		arr[i] = float64(sum) + float64(i)/1000
		m[arr[i]] = i
	}

	sort.Float64s(arr)

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		index := m[arr[i]]
		ans[i] = index
	}

	return ans
}

func main() {

}
