package main

import "fmt"

func reverseVowels(s string) string {

	m := map[uint8]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	l := 0
	r := len(s) - 1
	for l < r {
		_, ok1 := m[s[l]]
		_, ok2 := m[s[r]]
		if ok1 && ok2 {
			temp := s[0:l] + string(s[r]) + s[l+1:r] + string(s[l])
			if r < len(s) {
				temp += s[r+1:]
			}
			s = temp
			r--
			l++
		}
		if !ok1 {
			l++
		}
		if !ok2 {
			r--
		}
	}

	return s
}

func main() {

	//s := "hello"
	s := "leetcode"

	res := reverseVowels(s)

	fmt.Printf("res: %v\n", res)
}
