package main

import "fmt"

// func minimumSteps(s string) int64 {
//
//		n := len(s)
//		r := n - 1
//		r2 := r
//
//		// 01234
//		// 11000
//		// r=4
//		skip := 0
//		var ans int64
//		for r >= 0 && r2 >= 0 {
//
//			// 右边移动到'0'位置上
//			for s[r] == '1' && r >= 0 {
//				r--
//			}
//			if r < 0 {
//				return ans
//			}
//
//			for s[r2] == '0' && r2 >= 0 {
//				r2--
//			}
//			if r2 < 0 {
//				return ans
//			}
//
//			ans += int64(r - r2)
//			skip++
//			r--
//			r2--
//		}
//
//		return ans
//	}
func minimumSteps(s string) int64 {

	n := len(s)
	l := 0
	r := n - 1

	var ans int64
	for l < r {

		for r >= 0 && s[r] == '1' {
			r--
		}
		for l <= r && s[l] == '0' {
			l++
		}
		if r <= l {
			break
		}
		ans += int64(r - l)
		r--
		l++
	}

	return ans
}

func main() {
	// 6
	//s := "11000"
	s := "0111"
	// 4
	//s := "11001"
	res := minimumSteps(s)

	fmt.Printf("res: %v\n", res)
}
