package main

import "fmt"

func average(salary []int) float64 {

	minS := salary[0]
	maxS := minS
	sum := minS

	n := len(salary)
	for i := 1; i < n; i++ {
		if salary[i] > maxS {
			maxS = salary[i]
		}
		if salary[i] < minS {
			minS = salary[i]
		}
		sum += salary[i]
	}

	return float64(sum-minS-maxS) / float64(n-2)
}

func main() {
	salary := []int{4000, 3000, 1000, 2000}
	res := average(salary)

	fmt.Printf("res: %v\n", res)
}
