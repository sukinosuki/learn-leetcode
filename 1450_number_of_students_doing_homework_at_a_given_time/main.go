package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {

	cnt := 0

	n := len(startTime)
	for i := 0; i < n; i++ {

		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			cnt++
		}
	}

	return cnt
}

func main() {

}
