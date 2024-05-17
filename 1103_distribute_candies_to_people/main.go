package main

import "fmt"

func distributeCandies(candies int, num_people int) []int {

	ans := make([]int, num_people)

	n := 1
	for candies > 0 {

		for i := 0; i < num_people && candies > 0; i++ {
			if candies >= n {
				ans[i] += n
			} else {
				ans[i] += candies
			}
			candies -= n
			n++

		}
	}

	return ans
}

func main() {

	// [1 2 3 4]
	//candies := 7
	//num_people := 4

	// [5 2 3]
	candies := 10
	num_people := 3
	res := distributeCandies(candies, num_people)

	fmt.Printf("res: %v\n", res)
}
