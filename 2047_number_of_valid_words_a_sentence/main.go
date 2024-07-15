package main

import "fmt"

func countValidWords(sentence string) int {

	ans := 0

	l := 0
	n := len(sentence)
	for l < n {
		for l < n && sentence[l] == ' ' {
			l++
		}
		if l == n {
			return ans
		}
		flag := true

		start := l
		slashCount := 0
		for l < n && sentence[l] != ' ' {
			if !isLetter(sentence[l]) {
				if sentence[l] == '-' {
					if slashCount > 0 {
						flag = false
					}
					slashCount++
					if l-start == 0 || l == n || !isLetter(sentence[l+1]) {
						flag = false
					}
					l++
					continue
				}

				if sentence[l] == ',' || sentence[l] == '!' || sentence[l] == '.' {
					if l != n-1 && sentence[l+1] != ' ' {
						flag = false
					}
					l++
					continue
				}

				flag = false
			}
			l++
		}

		if flag {
			ans++
		}
	}

	return ans
}
func isLetter(char uint8) bool {
	return char-'a' >= 0 && char-'a' <= 26
}

func main() {
	// 3
	//sentence := "cat and  dog"

	// 0
	//sentence := "!this  1-s b8"

	//sentence := "alice and  bob are playing stone-game10"
	//sentence := "he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."

	sentence := "a-b-c"
	res := countValidWords(sentence)

	fmt.Printf("res: %v\n", res)
}
