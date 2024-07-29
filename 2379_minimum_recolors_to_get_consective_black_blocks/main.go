package main

import "fmt"

func minimumRecolors(blocks string, k int) int {

	ans := k
	wCount := 0

	for i := 0; i < len(blocks); i++ {
		if i < k-1 {
			if blocks[i] == 'W' {
				wCount++
			}
			continue
		}

		if blocks[i] == 'W' {
			wCount++
		}

		ans = min(ans, wCount)

		if blocks[i-k+1] == 'W' {
			wCount--
		}
	}

	return ans
}

func main() {
	// 3
	blocks := "WBBWWBBWBW"
	k := 7
	//blocks := "WBWBBBW"
	//k := 2
	res := minimumRecolors(blocks, k)

	fmt.Printf("res: %v\n", res)
}
