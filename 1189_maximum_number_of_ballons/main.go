package main

import "fmt"

func maxNumberOfBalloons(text string) int {

	arr := make([]int, 26)
	for _, c := range text {
		arr[c-'a']++
	}

	count := 0
	balloon := "balloon"
	for true {
		flag := true
		for _, c := range balloon {
			if arr[c-'a'] <= 0 {
				flag = false
				break
			}
			arr[c-'a']--
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
