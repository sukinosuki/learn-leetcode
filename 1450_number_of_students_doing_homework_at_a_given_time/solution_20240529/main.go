package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) (ans int) {
	maxEndTime := 0
	for _, e := range endTime {
		maxEndTime = max(maxEndTime, e)
	}
	if queryTime > maxEndTime {
		return
	}
	cnt := make([]int, maxEndTime+2)
	for i, s := range startTime {
		cnt[s]++
		cnt[endTime[i]+1]--
	}
	for _, c := range cnt[:queryTime+1] {
		ans += c
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {

	//startTime := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	//endTime := []int{10, 10, 10, 10, 10, 10, 10, 10, 10}
	//queryTime := 5

	startTime := []int{1, 1, 1, 1}
	endTime := []int{1, 3, 2, 4}
	queryTime := 4

	res := busyStudent(startTime, endTime, queryTime)

	fmt.Printf("res: %v\n", res)
}
