package main

import "fmt"

func getMaximumGenerated(n int) int {

	if n == 0 {
		return 0
	}
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1

	i := 1
	maxNum := 1
	for 2*i <= n && 2*i+1 <= n {

		arr[i*2] = arr[i]
		arr[i*2+1] = arr[i] + arr[i+1]

		if arr[i*2+1] > maxNum {
			maxNum = arr[i*2+1]
		}
		i++
	}

	return maxNum
}

func main() {

	// 3
	//n:=7

	n := 2
	res := getMaximumGenerated(n)

	fmt.Printf("res: %v\n", res)
}
