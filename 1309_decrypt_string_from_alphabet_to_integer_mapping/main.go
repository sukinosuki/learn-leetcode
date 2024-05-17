package main

import "fmt"

func freqAlphabets(s string) string {

	var stack []byte
	i := 0
	n := len(s)
	for i < n {
		if s[i] == '#' {

			num1 := stack[len(stack)-1] - 96
			num2 := stack[len(stack)-2] - 96

			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = num1 + num2*10 + 96
		} else {
			stack = append(stack, s[i]+48)
		}

		i++
	}

	return string(stack)
}

// "a" = 97
//
//	1  = 49
func main() {

	//s := "01abcdefghij"
	////s := "01234"
	//
	//arr := []byte{97, 98, 99}
	//for i := range s {
	//	c := s[i]
	//	fmt.Printf("%v: %v\n", string(c), c)
	//}
	//
	//fmt.Printf("arr: %v\n", string(arr))

	// jkab
	s := "10#11#12"

	//acz
	//s := "1326#"
	res := freqAlphabets(s)

	fmt.Printf("res: %v\n", res)
}
