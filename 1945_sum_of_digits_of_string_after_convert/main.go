package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func getLucky(s string, k int) int {
//
//	sb := strings.Builder{}
//	for i := range s {
//
//		str := strconv.Itoa(int(s[i] - 'a' + 1))
//		sb.WriteString(str)
//	}
//	str := sb.String()
//	sum, _ := strconv.Atoi(str)
//
//	for k > 0 {
//
//		temp := 0
//		for sum > 0 {
//			temp += sum % 10
//			sum /= 10
//		}
//
//		sum = temp
//		k--
//	}
//
//	return sum
//
// }
func getLucky(s string, k int) int {

	sb := strings.Builder{}
	for i := range s {

		str := strconv.Itoa(int(s[i] - 'a' + 1))
		sb.WriteString(str)
	}
	str := sb.String()

	for k > 0 {

		sum := 0

		for i := range str {
			sum += int(str[i] - '0')
		}

		str = strconv.Itoa(sum)

		k--
	}

	ans, _ := strconv.Atoi(str)
	return ans
}

func main() {

	// 36
	s := "iiii"
	k := 1
	// 8
	//s := "zbax"
	//k := 2
	// 8
	//s := "fleyctuuajsr"
	//k := 5
	res := getLucky(s, k)

	fmt.Printf("res: %v\n", res)
}
