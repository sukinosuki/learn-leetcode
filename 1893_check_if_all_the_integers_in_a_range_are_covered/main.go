package main

func isCovered(ranges [][]int, left int, right int) bool {

	cnt := make([]int, 51)

	for _, arr := range ranges {

		for i := arr[0]; i <= arr[1]; i++ {
			cnt[i] = 1
		}
	}

	for i := left; i <= right; i++ {
		if cnt[i] == 0 {
			return false
		}
	}

	return true
}

func main() {

}
