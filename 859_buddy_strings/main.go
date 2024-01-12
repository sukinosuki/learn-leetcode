package main

import "fmt"

func buddyStrings(s string, goal string) bool {

	m := len(s)
	n := len(goal)

	if m != n {
		return false
	}

	arr := make([]int, 0)

	for i := 0; i < m; i++ {
		if s[i] != goal[i] {
			if len(arr) == 2 {
				return false
			}
			arr = append(arr, i)
		}
	}

	if len(arr) == 2 {
		return s[arr[0]] == goal[arr[1]] && s[arr[1]] == goal[arr[0]]
	}

	if len(arr) == 0 {
		m := make(map[uint8]int)
		for i := range s {
			m[s[i]]++
			if m[s[i]] >= 2 {
				return true
			}
		}
	}

	return false
}

func main() {
	//s := "ab"
	//goal := "ba"

	//s := "ab"
	//goal := "ab"

	//s := "aa"
	//goal := "aa"

	s := "abcaa"
	goal := "abcbb"
	res := buddyStrings(s, goal)
	fmt.Printf("res: %v\n", res)
}
