package main

import "fmt"

func countBalls(lowLimit int, highLimit int) int {

	ans := 0
	cnt := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {

		sum := 0
		num := i
		for num >= 10 {
			sum += num % 10
			num = num / 10
		}
		sum += num
		cnt[sum]++
		ans = max(ans, cnt[sum])
	}

	return ans
}

func main() {
	// 2
	//lowLimit := 1
	//highLimit := 10

	// 2
	//lowLimit := 5
	//highLimit := 15
	// 2
	//lowLimit := 19
	//highLimit := 28

	lowLimit := 11
	highLimit := 104

	res := countBalls(lowLimit, highLimit)

	fmt.Printf("res: %v\n", res)
}
