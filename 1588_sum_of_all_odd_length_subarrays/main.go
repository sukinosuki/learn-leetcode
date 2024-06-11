package main

import "fmt"

func sumOddLengthSubarrays(arr []int) int {

	// [1,4,2,5,3]
	// [0,1,5,7,12,15]
	n := len(arr)
	sumArr := make([]int, n+1)

	for i := range arr {
		sumArr[i+1] += sumArr[i] + arr[i]
	}

	sum := sumArr[n]
	step := 3
	for step <= n {

		count := step
		for count <= n {

			sum += sumArr[count] - sumArr[count-step]
			count++
		}
		step += 2
	}

	return sum
}

func main() {
	// 58
	//arr := []int{1, 4, 2, 5, 3}

	//3
	//arr := []int{1, 2}

	arr := []int{10, 11, 12}
	res := sumOddLengthSubarrays(arr)

	fmt.Printf("res: %v\n", res)
}
