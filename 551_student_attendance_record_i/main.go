package main

import "fmt"

func checkRecord(s string) bool {

	absentCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			absentCount++
		}
		if absentCount >= 2 {
			return false
		}

		if i >= 2 {
			if s[i] == 'L' && s[i-1] == 'L' && s[i-2] == 'L' {
				return false
			}
		}
	}

	return true
}

func main() {
	//s := "AA"
	//s := "PPALLLP"
	s := "ALLAPPL"

	hasReward := checkRecord(s)

	fmt.Printf("hasReward: %v\n", hasReward)
}
