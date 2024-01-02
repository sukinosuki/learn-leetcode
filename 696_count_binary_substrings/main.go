package main

import "fmt"

func countBinarySubstrings(s string) int {

	//stack := []uint8{}
	//ans := 0
	//for i := 0; i < len(s); i++ {
	//	stack = []uint8{s[i]}
	//
	//	for j := i + 1; j < len(s); j++ {
	//		if s[i] == s[j] {
	//			stack = append(stack, s[j])
	//		} else {
	//			stack = stack[:len(stack)-1]
	//		}
	//
	//		if len(stack) == 0 {
	//			ans++
	//			break
	//		}
	//	}
	//
	//}
	//
	//return ans
}

func main() {
	str := "00110011"
	count := countBinarySubstrings(str)

	fmt.Printf("count: %v\n", count)
}
