package main

import "fmt"

func minTimeToType(word string) int {

	ans := 0

	var prev uint8 = 'a'
	for i := range word {

		n := int(word[i]) - int(prev)

		if n < 0 {
			n = -n
		}
		if n < 13 {
			ans += n
		} else {
			ans += 26 - n
		}

		ans++

		prev = word[i]
	}

	return ans
}

func main() {

	// 5
	word := "abc"
	// 7
	//word := "bza"
	res := minTimeToType(word)

	fmt.Printf("res: %v\n", res)
}
