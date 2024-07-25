package main

import "fmt"

func digitCount(num string) bool {

	cnt := make(map[int]int, 11)
	for i := range num {
		cnt[int(num[i]-'0')]++
	}

	for i := range num {
		if cnt[i] != int(num[i]-'0') {
			return false
		}
	}

	return true
}

func main() {
	// true
	//num := "1210"
	num := "030"
	res := digitCount(num)

	fmt.Printf("res: %v\n", res)
}
