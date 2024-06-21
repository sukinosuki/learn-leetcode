package main

import "fmt"

func temperatureTrend(temperatureA []int, temperatureB []int) int {

	ans := 0

	l1 := 1

	n := len(temperatureA)

	temp := ans
	for l1 < n {

		if (temperatureA[l1] > temperatureA[l1-1] && temperatureB[l1] > temperatureB[l1-1]) || (temperatureA[l1] < temperatureA[l1-1] && temperatureB[l1] < temperatureB[l1-1]) || (temperatureA[l1] == temperatureA[l1-1] && temperatureB[l1] == temperatureB[l1-1]) {
			temp++
		} else {
			temp = 0
		}

		ans = max(ans, temp)
		l1++
	}

	return ans
}

func main() {
	// 2
	//a := []int{21, 18, 18, 18, 31}
	//b := []int{34, 32, 16, 16, 17}
	a := []int{5, 10, 16, -6, 15, 11, 3}
	b := []int{16, 22, 23, 23, 25, 3, -16}
	res := temperatureTrend(a, b)

	fmt.Printf("res: %v\n", res)
}
