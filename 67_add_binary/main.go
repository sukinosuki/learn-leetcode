package main

import (
	"fmt"
	"strconv"
)

// 输入：a = "1010", b = "1011"
// 输出："10101"
func addBinary(a string, b string) string {

	l1 := len(a) - 1
	l2 := len(b) - 1

	carry := 0
	res := ""
	for l1 >= 0 || l2 >= 0 {
		num1 := 0
		num2 := 0
		if l1 >= 0 {
			num1, _ = strconv.Atoi(string(a[l1]))
			l1--
		}

		if l2 >= 0 {
			num2, _ = strconv.Atoi(string(b[l2]))
			l2--
		}

		sum := num1 + num2 + carry
		carry = sum % 2

		res = strconv.Itoa(sum%2) + res

		// 判断提前退出
		if carry == 0 && l1 < 0 && l2 > 0 {
			return b[:l2+1] + res
		}
		if carry == 0 && l2 < 0 && l1 > 0 {
			return a[:l1+1] + res
		}
	}

	if carry > 0 {
		return "1" + res
	}

	return res
}

func main() {
	//a := "1010"
	//b := "1011"

	//a := "1111"
	//b := "11111111"

	a := "100"
	//b := "110010"

	sub := a[1] - '0'

	fmt.Printf("sub: %d\n", sub)

	//res := addBinary(a, b)
	//fmt.Printf("res: %s\n", res)
}
