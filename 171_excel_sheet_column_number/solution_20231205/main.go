package main

import "fmt"

func titleToNumber(columnTitle string) int {

	number := 0
	for i, multiple := len(columnTitle)-1, 1; i >= 0; i-- {
		k := columnTitle[i] - 'A' + 1
		number += int(k) * multiple

		multiple *= 26
	}

	return number
}

func main() {
	//str := "A"
	//str := "AB"
	str := "ZY" // 701

	number := titleToNumber(str)

	fmt.Printf("number: %v\n", number)
}
