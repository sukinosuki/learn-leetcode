package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make(map[uint8]int)
	m2 := make(map[uint8]int)
	l := 0
	r := 0
	for index := range s {
		if s[index] == t[index] {
			continue
		}
		if value, ok := m1[t[index]]; ok && value > 0 {
			m1[t[index]] -= 1
			l--
		} else {
			l++
			m2[t[index]] += 1
		}

		if value, ok := m2[s[index]]; ok && value > 0 {
			m2[s[index]] -= 1
			r--
		} else {
			r++
			m1[s[index]] += 1
		}
	}

	//return l == 0 && r == 0
	return l+r == 0
}

func main() {
	//s := "anagram"
	//t := "nagaram"

	//s := "dgqztusjuu"
	//t := "dqugjzutsu"
	s := "yqhbicoumu"
	t := "ouiuycbmqh"

	isValid := isAnagram(s, t)

	fmt.Printf("isValid: %v\n", isValid)
}
