package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {

	sb := &strings.Builder{}
	maxLen := len(num1)
	if maxLen < len(num2) {
		maxLen = len(num2)
	}

	arr := make([]int, maxLen)
	l1 := len(num1) - 1
	l2 := len(num2) - 1
	carry := 0
	for maxLen > 0 {
		number1 := 0
		number2 := 0
		if l1 >= 0 {
			//number1, _ = strconv.Atoi(string(num1[l1]))
			number1 = int(num1[l1] - '0')
		}
		if l2 >= 0 {
			number2, _ = strconv.Atoi(string(num2[l2]))
		}

		carry = number1 + number2 + carry
		arr[maxLen-1] = carry % 10
		carry /= 10
		l1--
		maxLen--
		l2--
	}
	if carry > 0 {
		sb.WriteString(strconv.Itoa(carry))
	}
	for _, value := range arr {
		sb.WriteString(strconv.Itoa(value))
	}

	return sb.String()
}

func main() {
	//num1 := "11"
	//num2 := "123"
	num1 := "456"
	num2 := "77"
	sum := addStrings(num1, num2)
	fmt.Printf("sum: %v\n", sum)
}
