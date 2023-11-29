package main

import "fmt"

func strStr(haystack string, needle string) int {

	for i := 0; i < len(haystack); i++ {
		if len(haystack)-i < len(needle) {
			return -1
		}

		p1 := 0
		for p1 < len(needle) {
			if haystack[i+p1] != needle[p1] {
				break
			}
			p1++

			if p1 == len(needle) {
				return i
			}
		}
	}

	return -1
}

func main() {
	//haystack := "abcdefg"
	//needle := "cd"

	haystack := "sadbutsad"
	needle := "sad"

	//haystack := "leetcode"
	//needle := "leeto"

	//haystack := "aaa"
	//needle := "aaaa"

	index := strStr(haystack, needle)
	fmt.Printf("index: %d\n", index)
}
