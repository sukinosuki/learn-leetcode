package main

import (
	"fmt"
)

func findLUSlength(strs []string) int {

	ans := -1
	m := make(map[int]struct{})
	for i := 0; i < len(strs); i++ {

		if _, ok := m[i]; ok {
			continue
		}

		j := i + 1
		flag := true
		for ; j < len(strs); j++ {
			if check(strs[i], strs[j]) {
				if len(strs[i]) == len(strs[j]) {
					flag = false
					m[j] = struct{}{}
					// i 是 j的子字符串
				} else if len(strs[i]) < len(strs[j]) {
					flag = false
					// j是i的子字符串
				} else {
					m[j] = struct{}{}
				}
			}
		}

		if flag {
			ans = max(len(strs[i]), ans)
		}
	}

	return ans
}

func check(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}

	longStr := s1
	shortStr := s2
	if len(s1) < len(s2) {
		longStr, shortStr = s2, s1
	}

	l1 := 0
	l2 := 0
	for l1 < len(shortStr) && l2 < len(longStr) {

		if shortStr[l1] != longStr[l2] {
			l2++
			continue
		}

		l1++
		l2++

	}

	return l1 == len(shortStr)
}

func main() {
	// 3
	//strs := []string{"aba", "cdc", "eae"}
	//7
	//strs := []string{"aabbcc", "aabbcc", "bc", "bcc", "aabbccc"}
	// 2
	//strs := []string{"aabbcc", "aabbcc", "cb"}
	strs := []string{"aabbcc", "aabbcc", "cb", "abc"}
	res := findLUSlength(strs)

	fmt.Printf("res: %v\n", res)
}
