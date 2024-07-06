package main

import "fmt"

func largestOddNumber(num string) string {

	n := len(num)

	for j := n - 1; j >= 0; j-- {

		if (num[j]-'0')%2 == 1 {
			return num[:j+1]
		}
	}

	return ""
}

func main() {

	// 5
	//num := "52"
	num := "4206"
	res := largestOddNumber(num)

	fmt.Printf("res: %v\n", res)
}
