package main

import "fmt"

func distanceBetweenBusStops(distance []int, start int, destination int) int {

	//  1 3 6 10
	// [1,2,3,4]
	if start > destination {
		start, destination = destination, start
	}
	for i := 1; i < len(distance); i++ {
		distance[i] = distance[i-1] + distance[i]
	}

	distance2 := distance[destination-1]
	if start > 0 {
		distance2 -= distance[start-1]
	}
	distance1 := distance[len(distance)-1] - distance2

	if distance1 < distance2 {
		return distance1
	}

	return distance2
}

func main() {

	// 1 9
	//distance := []int{1, 2, 3, 4}
	//start := 0
	//destination := 1

	// 3 7
	//distance := []int{1, 2, 3, 4}
	//start := 0
	//destination := 2

	distance := []int{1, 2, 3, 4}
	start := 2
	destination := 3

	res := distanceBetweenBusStops(distance, start, destination)

	fmt.Printf("res: %v\n", res)
}
