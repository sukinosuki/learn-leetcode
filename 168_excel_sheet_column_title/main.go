package main

import "fmt"

func convertToTitle(columnNumber int) string {

	s := ""

	n := columnNumber

	for n > 26 {
		delivery := n % 26
		n = n / 26
		if delivery == 0 && n > 0 {
			delivery = 26
			n -= 1
		}
		if delivery > 0 {
			char := string(uint8(delivery + 64))
			s = char + s
		}
	}

	return string(uint8(n+64)) + s
}

// A: 65
// Z: 90
func main() {
	//s := "ABCDZ"
	//
	//for index := range s {
	//
	//	c := s[index]
	//
	//	fmt.Printf("c %v\n", c)
	//}
	//
	//a := uint8(65)
	//fmt.Println("a ", string(a))
	//columnNumber := 1
	columnNumber := 28 // AB
	//columnNumber := 701
	//columnNumber := 2147483647
	//columnNumber := 52 // AZ
	str := convertToTitle(columnNumber)

	fmt.Printf("str: %v\n", str)
}
