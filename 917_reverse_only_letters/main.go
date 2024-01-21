package main

import "fmt"

func reverseOnlyLetters(s string) string {

	n := len(s)
	l := 0
	r := n - 1
	arr := make([]byte, n)

	for l < n {

		// 是字母
		if (s[l] >= 97 && s[l] <= 122) || (s[l] >= 65 && s[l] <= 90) {
			if l <= r {
				for !((s[r] >= 97 && s[r] <= 122) || (s[r] >= 65 && s[r] <= 90)) {
					r--
				}
				if r >= l {
					arr[l], arr[r] = s[r], s[l]
				}
				r--
			}
		} else {
			arr[l] = s[l]
			l++
			continue
		}

		l++
	}

	return string(arr)
}

// s := "azAZ"
// a 97
// z 122
// A 65
// Z 90
func main() {

	//s := "ab-cd"
	//s := "a-bC-dEf-ghIj"
	s := "a"
	//s := "Test1ng-Leet=code-Q!"
	res := reverseOnlyLetters(s)

	fmt.Printf("res: %v\n", res)
}
