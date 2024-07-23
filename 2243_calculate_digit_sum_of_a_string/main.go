package main

import (
	"fmt"
	"strconv"
)

func digitSum(s string, k int) string {

	arr := make([]int, len(s))

	for i := range s {
		arr[i] = int(s[i] - '0')
	}

	sum := 0
	for len(arr) > k {
		temp := arr[:]

		arr = []int{}
		for i := 0; i < len(temp); i++ {

			sum += temp[i]
			if ((i+1) >= k && (i+1)%k == 0) || i == len(temp)-1 {

				if sum >= 10 {
					var subArr []int
					for sum > 0 {
						subArr = append([]int{sum % 10}, subArr...)
						sum /= 10
					}
					arr = append(arr, subArr...)
				} else {
					arr = append(arr, sum)
				}

				sum = 0
			}
		}
	}

	ans := ""
	for i := range arr {
		ans += strconv.Itoa(arr[i])
	}

	return ans
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
