package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {

	ans := 0

	for {
		for i := range tickets {
			if tickets[i] > 0 {
				tickets[i]--
				ans++
				if k == i && tickets[i] == 0 {
					return ans
				}
			}
		}
	}
}

func main() {
	// 6
	//tickets := []int{2, 3, 2}
	//k := 2
	// 8
	tickets := []int{5, 1, 1, 1}
	k := 0
	res := timeRequiredToBuy(tickets, k)
	fmt.Printf("res: %v\n", res)
}
