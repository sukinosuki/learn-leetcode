package main

import "fmt"

func threeConsecutiveOdds(arr []int) bool {

	n := len(arr)

	i := 0
	for i < n-2 {

		if arr[i]%2 == 0 {
			i++
			continue
		}
		if arr[i+1]%2 == 0 {
			i += 2
			continue
		}

		if arr[i+2]%2 == 0 {
			i += 3
			continue
		}

		return true
	}

	return false
}

func main() {
	// false
	//arr := []int{2, 6, 4, 1}
	arr := []int{1, 2, 34, 3, 4, 5, 7, 23, 12}
	res := threeConsecutiveOdds(arr)

	fmt.Printf("res: %v\n", res)
}
