package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	var cnt map[byte]int

	n := len(s)

	maxCount := 0
	for i := 0; i < n; i++ {
		cnt = map[byte]int{}

		cnt[s[i]]++
		count := 1
		for j := i + 1; j < n; j++ {
			if cnt[s[j]] == 0 {
				count++
				cnt[s[j]]++
			} else {
				break
			}
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)

	fmt.Printf("res: %v\n", res)
}
