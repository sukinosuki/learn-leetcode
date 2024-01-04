package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {

	for _, v := range letters {

		if v > target {
			return v
		}
	}

	return letters[0]
}

func main() {
	letters := []byte{'c', 'f', 'j'}
	res := nextGreatestLetter(letters, 'a')

	fmt.Printf("res: %v\n", string(res))
}
