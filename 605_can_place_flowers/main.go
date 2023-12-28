package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {

	l1 := 0
	count := 0

	for l1 < len(flowerbed) {
		sum := flowerbed[l1]
		if l1 > 0 {
			sum += flowerbed[l1-1]
		}
		if l1 < len(flowerbed)-1 {
			sum += flowerbed[l1+1]
		}

		if sum == 0 {
			flowerbed[l1] = 1
			count++
		}

		l1++
	}

	return count >= n
}

func main() {
	nums := []int{1, 0, 0, 0, 1}
	res := canPlaceFlowers(nums, 1)

	fmt.Printf("res: %v\n", res)
}
