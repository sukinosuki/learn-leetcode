package main

import (
	"fmt"
)

func mostVisited(n int, rounds []int) []int {

	cnt := make([]int, n+1)

	for i := 0; i < len(rounds)-1; i++ {

		if rounds[i] > rounds[i+1] {

			count := rounds[i] + 1

			for count <= n {
				cnt[count]++
				count++
			}

			count = 1
			for count <= rounds[i+1] {
				cnt[count]++
				count++
			}
		} else {
			count := rounds[i] + 1
			for count <= rounds[i+1] {
				cnt[count]++
				count++
			}
		}
	}
	cnt[rounds[0]]++

	var ans []int

	maxCount := 0

	for i := 1; i <= n; i++ {
		if cnt[i] == maxCount {
			ans = append(ans, i)
		} else if cnt[i] > maxCount {
			ans = []int{i}
			maxCount = cnt[i]
		}
	}
	return ans
}

func main() {
	//n := 4
	//  [1 2]
	//rounds := []int{1, 3, 1, 2}

	// [2]
	//n := 2
	//rounds := []int{2, 1, 2, 1, 2, 1, 2, 1, 2}
	n := 7
	rounds := []int{1, 3, 5, 7}
	res := mostVisited(n, rounds)

	fmt.Printf("res: %v\n", res)
}
