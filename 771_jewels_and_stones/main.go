package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {

	ans := 0
	arr := make([]int, 58)
	for _, c := range jewels {
		arr[c-'A'] = 1
	}

	for _, c := range stones {
		if arr[c-'A'] > 0 {
			ans++
		}
	}

	return ans
}

// s := "azAZ"
// a 97
// z 122
// A 65
// Z 90
func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	res := numJewelsInStones(jewels, stones)

	fmt.Printf("res: %v\n", res)
}
