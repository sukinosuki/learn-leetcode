package main

import "fmt"

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {

	m1 := make(map[int]int)
	m2 := make(map[int]int)
	aliceSum := 0
	bobSum := 0
	for _, num := range aliceSizes {
		aliceSum += num
		m1[num] = 1
	}

	for _, num := range bobSizes {
		bobSum += num
		m2[num] = 1
	}

	if aliceSum > bobSum {
		quotient := (aliceSum - bobSum) / 2

		for num := range m1 {
			if m2[num-quotient] > 0 {

				return []int{num, num - quotient}
			}
		}
	} else {
		quotient := (bobSum - aliceSum) / 2
		for num := range m2 {
			if m1[num-quotient] > 0 {
				return []int{num - quotient, num}
			}
		}
	}

	return nil
}

func main() {
	//aliceSizes := []int{1, 2, 5}
	//bobSizes := []int{2, 4}

	aliceSizes := []int{1, 1}
	bobSizes := []int{2, 2}
	res := fairCandySwap(aliceSizes, bobSizes)

	fmt.Printf("res: %v\n", res)
}
