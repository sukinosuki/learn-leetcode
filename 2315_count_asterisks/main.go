package main

import "fmt"

func countAsterisks(s string) int {

	ans := 0

	for i := 0; i < len(s); i++ {

		if s[i] == '*' {
			ans++
			continue
		}

		if s[i] == '|' {
			i++
			for i < len(s) && s[i] != '|' {
				i++
			}
		}

	}

	return ans
}

func main() {
	s := "l|*e*et|c**o|*de|"
	res := countAsterisks(s)

	fmt.Printf("res: %v\n", res)
}
