package main

import "fmt"

func removeDigit(number string, digit byte) string {

	index := -1
	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			index = i
		}

		if number[i] == digit && i+1 < len(number) && number[i+1] > digit {
			break
		}
	}

	return number[:index] + number[index+1:]

}

func main() {
	// 12
	numbers := "123"
	digit := '3'
	// 231
	//numbers := "1231"
	//digit := '1'

	//numbers := "551"
	//digit := '5'

	//  3965
	//numbers := "83965"
	//digit := '8'
	res := removeDigit(numbers, byte(digit))

	fmt.Printf("res: %v\n", res)
}
