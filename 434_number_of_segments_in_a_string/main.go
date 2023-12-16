package main

import "fmt"

func countSegments(s string) int {

	n := len(s)
	l1 := 0
	l2 := 0
	count := 0
	for l1 < n {
		if l2 < n && s[l1] == ' ' && s[l2] == ' ' {
			l1++
			l2++
			continue
		}
		if l2 >= len(s)-1 || s[l2] == ' ' {
			l2++
			l1 = l2
			count++
			continue
		}

		l2++
	}

	return count
}

func main() {
	s := "Hello, my name is John"
	count := countSegments(s)

	fmt.Printf("count: %v\n", count)
}
