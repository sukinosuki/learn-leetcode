package main

import (
	"fmt"
)

// 复制解
func thousandSeparator(n int) string {

	ans := ""
	for n >= 1000 {
		ans = fmt.Sprintf(".%03d%s", n%1000, ans)
		n /= 1000
	}
	ans = fmt.Sprintf("%d%s", n, ans)
	return ans
}

func main() {
	//n := 987
	//n := 123456789
	n := 123456789
	res := thousandSeparator(n)

	fmt.Printf("res: %v\n", res)

	fmt.Sprintf(".%03d", 123)
}
