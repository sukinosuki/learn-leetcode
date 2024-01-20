package main

import "fmt"

func splitWordsBySeparator(words []string, separator byte) []string {
	ans := []string{}

	for _, str := range words {

		n := len(str)

		l1 := 0
		l2 := 0
		for l2 < n {
			if str[l1] == separator {
				l1++
				l2++
				continue
			}

			if str[l2] == separator {
				ans = append(ans, str[l1:l2])
				l2++
				l1 = l2
				continue
			}
			if l2 == n-1 {
				ans = append(ans, str[l1:])
			}

			l2++
		}
	}

	return ans
}

func main() {
	words := []string{"one.two.three", "four.five", "six"}
	res := splitWordsBySeparator(words, '.')
	//words := []string{"$easy$", "$problem$"}
	//res := splitWordsBySeparator(words, '$')

	fmt.Printf("res: %v\n", res)
}
