package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {

	total := 0
	prev := 0
	for index, time := range timeSeries {
		total += duration
		if index != 0 && time-prev < duration {
			total += time - prev - duration
		}
		prev = time
	}

	return total
}

func main() {
	timeSeries := []int{1, 4} // 4
	duration := 2

	//timeSeries := []int{1, 2} // 3
	//duration := 2

	//timeSeries := []int{1, 2, 3, 4, 5} // 9
	//duration := 5

	total := findPoisonedDuration(timeSeries, duration)
	fmt.Printf("total: %v\n", total)

}
