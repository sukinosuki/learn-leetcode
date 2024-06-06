package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {

	sort.Ints(arr)
	dis := arr[1] - arr[0]
	n := len(arr)
	for i := 2; i < n; i++ {
		if arr[i]-arr[i-1] != dis {
			return false
		}
	}

	return true
}

func main() {

}
