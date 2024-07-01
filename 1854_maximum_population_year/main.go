package main

import "fmt"

func maximumPopulation(logs [][]int) int {

	arr := make([]int, 101)

	for i := range logs {
		arr[logs[i][0]-1950]++
		arr[logs[i][1]-1950]--
	}

	ans := 0
	maxNum := 0

	num := 0
	for i := range arr {
		num += arr[i]
		if num > maxNum {
			maxNum = num

			ans = 1950 + i
		}
	}

	return ans
}

func main() {
	// 1993
	//logs := [][]int{
	//	{1993, 1999},
	//	{2000, 2010},
	//}
	logs := [][]int{
		{1950, 1961},
		{1960, 1971},
		{1970, 1981},
	}
	res := maximumPopulation(logs)

	fmt.Printf("res: %v\n", res)
}
