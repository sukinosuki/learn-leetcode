package main

import "fmt"

func maxNumberOfBalloons(text string) int {

	m := make(map[uint8]int)
	for i := range text {
		m[text[i]]++
	}

	count := 0
	balloon := "balloon"
	for true {
		flag := true
		for i := range balloon {
			if size, ok := m[balloon[i]]; ok {
				if size == 0 {
					flag = false
					break
				}
				m[balloon[i]]--
			} else {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		count++
	}

	return count
}

func main() {
	text := "loonbalxballpoon"
	res := maxNumberOfBalloons(text)

	fmt.Printf("res: %v\n", res)
}
