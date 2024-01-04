package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {

	ans := []int{}
	for i := left; i <= right; i++ {
		if helper(i) {
			ans = append(ans, i)
		}
	}
	return ans
}

func helper(num int) bool {

	originNum := num
	for num > 0 {
		remainder := num % 10
		if remainder == 0 {
			return false
		}
		if originNum%remainder != 0 {
			return false
		}

		num /= 10
	}

	return true
}

func main() {
	res := selfDividingNumbers(1, 22)

	fmt.Printf("res: %v\n", res)
}
