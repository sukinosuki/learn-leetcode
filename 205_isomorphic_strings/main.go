package main

import "fmt"

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	m1 := make(map[uint8]int)
	m2 := make(map[uint8]int)

	for index := range s {
		char1 := s[index]
		char2 := t[index]

		char1Index, ok1 := m1[char1]
		char2Index, ok2 := m2[char2]
		if ok1 != ok2 {
			return false
		}
		if char1Index != char2Index {
			return false
		}

		m1[char1] = index
		m2[char2] = index
	}
	return true
}

func main() {
	s := "egg"
	t := "add"
	isValid := isIsomorphic(s, t)

	fmt.Printf("isValid: %v\n", isValid)
}
