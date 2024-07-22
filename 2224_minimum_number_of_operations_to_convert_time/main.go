package main

import (
	"fmt"
	"strconv"
)

func convertTime(current string, correct string) int {

	hour1, _ := strconv.Atoi(current[:2])
	hour2, _ := strconv.Atoi(correct[:2])

	minutes1, _ := strconv.Atoi(current[3:])
	minutes2, _ := strconv.Atoi(correct[3:])
	minutes1 += hour1 * 60
	minutes2 += hour2 * 60

	diff := minutes2 - minutes1
	ans := 0

	ans += diff / 60
	diff = diff % 60

	ans += diff / 15
	diff %= 15

	ans += diff / 5
	diff %= 5

	ans += diff

	return ans
}

func main() {
	// 3
	//current := "02:30"
	//correct := "04:35"
	// 1
	current := "11:00"
	correct := "11:01"
	res := convertTime(current, correct)

	fmt.Printf("res: %v\n", res)
}
