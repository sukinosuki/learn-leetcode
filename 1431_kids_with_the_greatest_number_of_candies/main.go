package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {

	maxNum := 0

	for i := range candies {
		if candies[i] > maxNum {
			maxNum = candies[i]
		}
	}

	ans := make([]bool, len(candies))
	for i := range candies {
		ans[i] = candies[i]+extraCandies >= maxNum
	}

	return ans
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3

	res := kidsWithCandies(candies, extraCandies)
	fmt.Printf("res: %v\n", res)

}
