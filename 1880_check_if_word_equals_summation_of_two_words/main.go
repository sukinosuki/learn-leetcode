package main

import "fmt"

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {

	return toNum(firstWord)+toNum(secondWord) == toNum(targetWord)
}

func toNum(s string) int {

	n := len(s)
	num := 0
	for i := 0; i < n; i++ {
		num *= 10
		num += int(s[i] - 'a')
	}

	return num
}

func main() {
	f := "acb"
	s := "cba"
	t := "cdb"
	res := isSumEqual(f, s, t)

	fmt.Printf("res: %v\n", res)
}
