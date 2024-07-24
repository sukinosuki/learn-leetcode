package main

import "fmt"

func percentageLetter(s string, letter byte) int {

	n := len(s)

	count := 0
	runeLetter := rune(letter)
	for _, c := range s {
		if c == runeLetter {
			count++
		}
	}
	ans := float64(count) / float64(n)

	ans *= 100

	return int(ans)
}

func main() {
	//s := "foobar"
	//letter := 'o'
	s := "vmvvvvvzrvvpvdvvvvyfvdvvvvpkvvbvvkvvfkvvvkvbvvnvvomvzvvvdvvvkvvvvvvvvvlvcvilaqvvhoevvlmvhvkvtgwfvvzy"
	letter := 'v'
	res := percentageLetter(s, byte(letter))
	fmt.Printf("res %v\n", res)
}
