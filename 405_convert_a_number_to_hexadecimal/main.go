package main

import (
	"fmt"
	"strconv"
)

func toHex(num int) string {
	ans := ""
	if num < 0 {
		num = 4294967295 + num + 1
	}

	for num > 15 {
		delivery := num % 16
		num /= 16

		ans = getHexByNum(delivery) + ans

	}
	ans = getHexByNum(num) + ans

	return ans
}

func getHexByNum(num int) string {
	switch num {
	case 10:
		return "a"
	case 11:
		return "b"
	case 12:
		return "c"
	case 13:
		return "d"
	case 14:
		return "e"
	case 15:
		return "f"
	}

	return strconv.Itoa(num)
}

func main() {
	//num := 255 // ff
	//num := 26 // 1a
	//num := 2555 // 1a
	num := 2555 // 1a

	res := toHex(num)

	fmt.Printf("res: %v\n", res)
}
