package main

import "fmt"

//func reverseStr(s string, k int) string {
//
//	n := len(s)
//	arr := make([]byte, n)
//
//	k1 := k
//	loop := 0
//	for i := 0; i < n; i++ {
//		//需要反转
//		if loop%2 == 0 && k > 1 {
//			index := loop*2 + k1 - 1
//			if index > n-1-k+k1 {
//				index = n - 1 - k + k1
//			}
//			arr[i] = s[index]
//			k1--
//			if k1 == 0 {
//				loop++
//			}
//		} else {
//			arr[i] = s[i]
//			k1++
//			if k1 == k {
//				loop++
//			}
//		}
//	}
//
//	return string(arr)
//}

func reverseStr(s string, k int) string {

	n := len(s)
	arr := make([]byte, 0)

	stack := make([]byte, 0)
	loop := 0
	for i := 0; i < n; i++ {
		if loop%2 == 0 {

			stack = append(stack, s[i])
			if len(stack) == k || i == n-1 {
				loop++
				for len(stack) > 0 {
					char := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					arr = append(arr, char)
				}
			}
		} else {
			arr = append(arr, s[i])
			if (i+1)%k == 0 {
				loop++
			}
		}
	}

	return string(arr)
}

// fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlbyhefdcqkmxwholhytmhaf
// fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi

func main() {

	s := "abcdefg"
	//s := "a"
	res := reverseStr(s, 1)
	//s := "abcdefg"
	//res := reverseStr(s, 8)
	fmt.Printf("res: %v\n", res)
	//str := "hello"
	//
	//arr := make([]byte, len(str))
	//fmt.Printf("arr length: %v\n", arr)
	//
	//for i := 0; i < len(str); i++ {
	//
	//	fmt.Printf("char: %v\n", str[i])
	//	//arr = append(arr, str[i])
	//	arr[i] = str[i]
	//}
	//
	//fmt.Printf("arr: %v\n", arr)
	//fmt.Printf("arr: %v\n", string(arr))
}
