package main

import "fmt"

func largestGoodInteger(num string) string {

	ans := ""

	for i := 0; i < len(num)-2; i++ {

		if num[i] == num[i+1] && num[i] == num[i+2] {
			if ans == "" || num[i] > ans[0] {
				ans = num[i : i+3]
				if ans[0] == '9' {
					return ans
				}
			}
			i += 2
		}
	}

	return ans
}

func main() {
	num := "677713333999"
	res := largestGoodInteger(num)

	fmt.Printf("res: %v\n", res)
}
