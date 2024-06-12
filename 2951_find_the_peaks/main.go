package main

import "fmt"

func findPeaks(mountain []int) []int {

	var ans []int
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {

	//mountain := []int{2, 4, 4}
	mountain := []int{1, 4, 3, 8, 5}
	res := findPeaks(mountain)

	fmt.Printf("res: %v\n", res)
}
