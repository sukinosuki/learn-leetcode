package main

import (
	"fmt"
)

func slowestKey(releaseTimes []int, keysPressed string) byte {

	maxTime := releaseTimes[0]
	n := len(releaseTimes)
	ans := keysPressed[0]
	for i := 1; i < n; i++ {

		sub := releaseTimes[i] - releaseTimes[i-1]
		if sub > maxTime || (sub == maxTime && keysPressed[i] > ans) {
			maxTime = sub
			ans = keysPressed[i]
		}
	}

	return ans
}

func main() {

	//releaseTimes := []int{9, 29, 49, 50}
	//keyPressed := "cbcd"
	releaseTimes := []int{23, 34, 43, 59, 62, 80, 83, 92, 97}
	keyPressed := "qgkzzihfc"
	res := slowestKey(releaseTimes, keyPressed)

	fmt.Printf("res: %v\n", string(res))
}
