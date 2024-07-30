package main

import (
	"fmt"
	"strconv"
)

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {

	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	maxArrive := arriveAlice
	minLeave := leaveAlice
	if maxArrive < arriveBob {
		maxArrive = arriveBob
	}
	if minLeave > leaveBob {
		minLeave = leaveBob
	}

	if minLeave < maxArrive {
		return 0
	}

	startM, _ := strconv.Atoi(maxArrive[:2])
	startD, _ := strconv.Atoi(maxArrive[3:])

	endM, _ := strconv.Atoi(minLeave[:2])
	endD, _ := strconv.Atoi(minLeave[3:])

	ans := 0
	// 8-17 > 8-18
	if startM < endM {
		ans += days[startM-1] - startD + 1
		startD = 1
		startM++
	}
	for startM < endM {
		ans += days[startM-1]
		startM++
	}

	ans += endD - startD + 1

	return ans
}

func main() {
	// 3
	//arriveAlice := "08-15"
	//leaveAlice := "08-18"
	//arriveBob := "08-16"
	//leaveBob := "08-19"

	arriveAlice := "10-01"
	leaveAlice := "10-31"
	arriveBob := "11-01"
	leaveBob := "12-31"
	res := countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob)
	fmt.Printf("res: %v\n", res)
}
