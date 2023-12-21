package main

import "fmt"

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i*2 <= n; i++ {
		if n%i == 0 {
			match := true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

func main() {
	s := "abababac"
	ok := repeatedSubstringPattern(s)

	fmt.Printf("ok: %v\n", ok)
}
