package main

import "fmt"

func bestHand(ranks []int, suits []byte) string {

	mask := 0
	if suits[0] == suits[1] && suits[0] == suits[2] && suits[0] == suits[3] && suits[0] == suits[4] {
		mask |= 1 << 5
	}

	cnt := make([]int, 14)

	for _, rank := range ranks {
		cnt[rank]++
		if cnt[rank] == 2 {
			mask |= 1 << 2
		}
		if cnt[rank] == 3 {
			mask |= 1 << 3
		}
	}

	if mask>>5&1 == 1 {
		return "Flush"
	}
	if mask>>3&1 == 1 {
		return "Three of a Kind"
	}
	if mask>>2&1 == 1 {
		return "Pair"
	}

	return "High Card"
}

func main() {
	ranks := []int{13, 2, 3, 1, 9}
	suits := []byte{'a', 'a', 'a', 'a', 'a'}

	res := bestHand(ranks, suits)

	fmt.Printf("res: %v\n", res)
}
