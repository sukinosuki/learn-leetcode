package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {

	//ans := []int{area, 1}

	mid := int(math.Sqrt(float64(area)))

	for mid > 0 {
		delivery := area / mid
		if delivery*mid == area {
			ans := []int{delivery, mid}
			//ans[0] = delivery
			//ans[1] = mid
			return ans
		}

		mid--
	}

	//return ans
	return nil
}

func main() {
	//area := 37
	area := 122122
	res := constructRectangle(area)

	fmt.Printf("res: %v\n", res)
}
