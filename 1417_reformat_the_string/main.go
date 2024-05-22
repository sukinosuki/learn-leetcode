package main

import "fmt"

func reformat(s string) string {

	n := len(s)
	if n == 1 {
		return s
	}
	var ans []uint8

	l1 := 0
	l2 := 0
	for l1 < n || l2 < n {
		if l1 < n {
			for s[l1] >= 'a' {
				l1++
				if l1 >= n {
					return ""
				}
			}
			ans = append(ans, s[l1])
			l1++
		}

		if l2 < n {
			for s[l2] < 'a' {
				l2++
				if l2 >= n {
					return ""
				}
			}
			ans = append(ans, s[l2])
			l2++
		}

	}

	if len(ans) != n {
		return ""
	}

	return string(ans)

}

func main() {

	// a0b1c2
	//s := "a0b1c2"

	//""
	//s := "leetcode"

	s := "covid2019"
	res := reformat(s)

	fmt.Printf("res: %v\n", res)
}
