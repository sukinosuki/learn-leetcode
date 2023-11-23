package main

import "fmt"

func getNumByRomanS(s uint8) int {
	switch s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}

	return 0
}

func romanToInt(s string) int {
	var prev uint8
	sum := 0
	for i := 0; i < len(s); i++ {
		romanS := s[i]
		var num int
		switch romanS {
		case 'I':
			num = 1
		case 'V':
			num = 5
		case 'X':
			num = 10
		case 'L':
			num = 50
		case 'C':
			num = 100
		case 'D':
			num = 500
		case 'M':
			num = 1000
		default:
			num = 0
		}

		sum += num

		if prev == 'I' {
			if romanS == 'V' || romanS == 'X' {
				sum -= 2
			}
		} else if prev == 'X' {
			if romanS == 'L' || romanS == 'C' {
				sum -= 20
			}
		} else if prev == 'C' {
			if romanS == 'D' || romanS == 'M' {
				sum -= 200
			}
		}
		prev = romanS
	}

	return sum
}

func main() {
	//s := "III"
	//s := "IV"
	//s := "IX"
	//s := "LVIII"
	s := "MCMXCIV"
	res := romanToInt(s)

	fmt.Printf("res: %d\n", res)
}
