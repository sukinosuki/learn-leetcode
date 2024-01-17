package main

import "fmt"

func uncommonFromSentences(s1 string, s2 string) []string {

	l1 := 0
	l2 := 0
	m1 := make(map[string]int)
	ans := make([]string, 0)
	s := s1 + " " + s2 + " "
	for l2 < len(s) {
		if l2 == len(s)-1 || s[l2] == ' ' {

			m1[s[l1:l2]]++
			l2++
			l1 = l2
		} else {
			l2++

		}
	}

	for key, value := range m1 {
		if value == 1 {
			ans = append(ans, key)
		}

	}

	return ans
}

func main() {
	s1 := "this apple is sweet"
	s2 := "this apple is sour"
	res := uncommonFromSentences(s1, s2)
	fmt.Printf("res: %v\n", res)
}
