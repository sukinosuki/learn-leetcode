package main

import "fmt"

func constructRectangle(area int) []int {

	r := area
	l := 0
	ans := []int{area, 1}
	for l <= r {
		mid := l + (r-l)/2

		delivery := area / mid
		if delivery*mid == area {
			if mid-delivery < ans[0]-ans[1] {
				ans[0] = mid
				ans[1] = delivery
			}

			continue
		}

		if delivery*mid < area {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}

func main() {
	res := constructRectangle(122122)

	fmt.Printf("res: %v\n", res)
}
