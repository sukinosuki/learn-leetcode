package main

import "fmt"

func countLargestGroup(n int) int {

	cnt := make(map[int]int)

	for i := 1; i <= n; i++ {

		sum := getNumSum(i)

		cnt[sum]++

	}

	maxV := 0
	count := 0
	for _, v := range cnt {
		if v > maxV {
			maxV = v
			count = 0
		}
		if v == maxV {
			count++
		}
	}

	return count
}

func getNumSum(num int) int {

	sum := 0
	for num > 0 {
		sum += num % 10

		num /= 10
	}

	return sum
}

func main() {

	// 4
	//n := 13

	// 2
	//n := 2

	// 6
	n := 15

	// 6
	//n := 46
	res := countLargestGroup(n)

	fmt.Printf("res: %v\n", res)
}
