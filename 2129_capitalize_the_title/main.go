package main

import "fmt"

func capitalizeTitle(title string) string {

	arr := []byte(title)

	l := 0

	n := len(arr)

	for l < n {
		start := l

		for l < n && arr[l] != ' ' {
			if arr[l] < 97 {
				arr[l] += 32
			}
			l++
		}
		if l-start > 2 {
			arr[start] -= 32
		}
		l++
	}

	return string(arr)
}

func main() {
	// [97 122 65 90 32]
	//title := "azAZ "
	// Capitalize The Title
	//title := "capiTalIze tHe titLe"
	title := "First leTTeR of EACH Word"
	res := capitalizeTitle(title)

	fmt.Printf("res: %v\n", res)
}
