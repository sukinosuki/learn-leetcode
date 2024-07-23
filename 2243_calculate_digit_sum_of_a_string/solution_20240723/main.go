package main

import (
	"fmt"
	"strconv"
)

func digitSum(s string, k int) string {

	sum := 0
	var temp string
	for len(s) > k {
		temp = s
		s = ""
		for i := range temp {
			sum += int(temp[i] - '0')
			if (i+1 >= k && (i+1)%k == 0) || i == len(temp)-1 {
				s += strconv.Itoa(sum)
				sum = 0
			}
		}
	}

	return s
}

func main() {
	// 135
	//s := "11111222223"
	//k := 3
	//s := "00000000"
	//k := 3
	s := "01234567890"
	k := 2
	res := digitSum(s, k)

	fmt.Printf("res: %v\n", res)
}
