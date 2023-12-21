package main

import "strings"

func repeatedSubstringPattern(s string) bool {

	s2 := s + s

	s2 = s2[1 : len(s2)-1]

	return strings.Contains(s2, s)
}

func main() {

}
