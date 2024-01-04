package main

import "fmt"

func isOneBitCharacter(bits []int) bool {

	l1 := 0
	n := len(bits)
	for l1 < n {

		if bits[l1] == 0 {
			l1++
			if l1 == n {
				return true
			}
			continue
		}
		if bits[l1] == 1 {
			l1 += 2
		}
	}

	return false
}

func main() {

	//bits := []int{0, 1, 0}
	bits := []int{0, 0, 0}
	//bits := []int{1, 1, 1, 0}
	//bits := []int{1, 0, 0}
	valid := isOneBitCharacter(bits)

	fmt.Printf("valid: %v\n", valid)
}
