package main

import (
	"fmt"
	"strings"
)

func ReformatNumber(number string) string {

	number = strings.ReplaceAll(number, "-", "")
	number = strings.ReplaceAll(number, " ", "")
	i := 0

	var ans []byte
	n := len(number)
	for i < n {

		remainLen := len(number[i:])
		if remainLen <= 4 {
			if remainLen == 3 {
				ans = append(ans, number[i:i+3]...)
				i += 3
				continue
			}

			if remainLen == 4 {
				ans = append(ans, number[i:i+2]...)
				ans = append(ans, '-')
				i += 2
				continue
			}

			ans = append(ans, number[i:i+2]...)
			i += 2

		} else {
			ans = append(ans, number[i:i+3]...)
			ans = append(ans, '-')
			i += 3
		}
	}

	return string(ans)
}

func main() {
	//  123-456
	//number := "1-23-45 6"
	// 123-45-67
	//number := "123 4-567"
	// 123-456-78
	//number := "123 4-5678"
	// 12
	//number := "12"
	// 175-229-353-94-75
	number := "--17-5 229 35-39475 "
	res := ReformatNumber(number)

	fmt.Printf("res: %s\n", res)
}
