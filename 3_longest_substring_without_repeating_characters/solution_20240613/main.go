package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	cnt := make(map[byte]int)
	ans := 0

	l1 := 0
	l2 := 1
	n := len(s)
	if n > 0 {
		cnt[s[l1]]++
		ans = 1
	}

	for l2 < n {

		for l2 < n && cnt[s[l2]] == 0 {
			cnt[s[l2]]++
			l2++
		}

		if l2-l1 > ans {
			ans = l2 - l1
		}

		cnt[s[l1]]--
		l1++
	}

	return ans
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)

	fmt.Printf("res: %v\n", res)
}
