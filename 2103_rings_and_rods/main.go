package main

import "fmt"

func countPoints(rings string) int {

	cnt := make([]int, 10)

	for i := 0; i < len(rings); i += 2 {

		color := rings[i]
		index := int(rings[i+1] - '0')

		switch color {
		case 'R':
			cnt[index] |= 1
		case 'G':
			cnt[index] |= 1 << 1
		case 'B':
			cnt[index] |= 1 << 2
		}
	}

	ans := 0
	for i := range cnt {
		if cnt[i] == 7 {
			ans++
		}
	}

	return ans
}

func main() {
	// 1
	//rings := "B0B6G0R6R0R6G9"
	rings := "B0R0G0R9R0B0G0"
	res := countPoints(rings)

	fmt.Printf("res: %v\n", res)
}
