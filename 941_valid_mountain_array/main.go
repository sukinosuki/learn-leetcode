package main

import "fmt"

func validMountainArray(arr []int) bool {

	isMounting := true
	if len(arr) < 2 {
		return false
	}

	if arr[1] <= arr[0] {
		return false
	}

	for i := 2; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		}

		if !isMounting {
			if arr[i] > arr[i-1] {
				return false
			}
		}

		if arr[i] < arr[i-1] {
			isMounting = false
		}
	}

	return !isMounting
}

func main() {
	//nums := []int{2, 1}
	//nums := []int{0, 2, 3, 4, 5, 2, 1, 0}// true
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // false
	//nums := []int{0, 2, 3, 3, 5, 2, 1, 0} // false
	res := validMountainArray(nums)

	fmt.Printf("nums: %v\n", res)
}
