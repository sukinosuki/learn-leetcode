package main

import "fmt"

func findLUSlength(a string, b string) int {

	//if len(a) < len(b) {
	//	a, b = b, a
	//}
	//
	//ans := len(a)
	//l1 := 0
	//l2 := 0
	//// a: aefawfawfawfaw
	//// b: aefawfeawfwafwaef
	//for l1 < len(a) && l2 < len(b) {
	//	if a[l1] != b[l2] {
	//		l1++
	//	} else {
	//		l1++
	//		l2++
	//		ans--
	//	}
	//}
	//if ans == 0 {
	//	ans--
	//}
	//
	//return ans
}

func main() {
	a := "abcd"
	b := "aed"
	length := findLUSlength(a, b)

	fmt.Printf("length: %v\n", length)
}
