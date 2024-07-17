package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {

	base := tickets[k] - 1

	ans := base*len(tickets) + k

	for i := range tickets {
		diff := tickets[i] - base
		if i < k {
			diff--
		}
		if diff < 0 {
			ans += diff
		}
	}

	return ans + 1
}

func main() {
	// 6
	tickets := []int{2, 3, 2}
	k := 2
	// 8
	//tickets := []int{5, 1, 1, 1}
	//k := 0
	//tickets := []int{84, 49, 5, 24, 70, 77, 87, 8}
	//k := 3
	res := timeRequiredToBuy(tickets, k)
	fmt.Printf("res: %v\n", res)
}
