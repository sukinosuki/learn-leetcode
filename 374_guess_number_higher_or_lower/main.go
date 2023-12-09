package main

import "fmt"

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int {

	if num == 6 {
		return 0
	}
	if num < 6 {
		return -1
	}

	return 1
}

func guessNumber(n int) int {
	l := 0
	r := n
	for l <= r {
		mid := l + (r-l)/2

		res := guess(mid)
		if res == 0 {
			return mid
		}
		if res == -1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func main() {

	res := guessNumber(10)

	fmt.Printf("res: %v\n", res)

}
