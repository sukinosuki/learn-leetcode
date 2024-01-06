package main

import "fmt"

func shortestCompletingWord(licensePlate string, words []string) string {

	return ""
}

// s := "azAZ"
// a 97
// z 122
// A 65
// Z 90
func main() {
	//s := "1s3 PSt"
	//words := []string{"step", "steps", "stripe", "stepple"}
	s := "1s3 456"
	words := []string{"looks", "pest", "stew", "show"}

	res := shortestCompletingWord(s, words)

	fmt.Printf("res: %v\n", res)
}
