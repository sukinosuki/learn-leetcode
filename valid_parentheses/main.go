package main

import "fmt"

// func isValid(s string) bool {
//
//		var stack []int32
//
//		m := map[int32]int32{'(': ')', '[': ']', '{': '}'}
//		for _, v := range s {
//
//			switch v {
//			case '(', '[', '{':
//				stack = append(stack, v)
//			case ')', ']', '}':
//				if len(stack) == 0 {
//					return false
//				}
//
//				stackLastValue := stack[len(stack)-1]
//
//				if v != m[stackLastValue] {
//					return false
//				}
//				stack = stack[:len(stack)-1]
//			}
//		}
//
//		return len(stack) == 0
//	}
func isValid(s string) bool {

	var stack []int32

	for _, v := range s {

		switch v {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')

		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}

			if v != stack[len(stack)-1] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
func main() {
	s := "()"
	//s := "()[]{}"
	//s := "(]"

	res := isValid(s)

	fmt.Printf("res: %v\n", res)
}
