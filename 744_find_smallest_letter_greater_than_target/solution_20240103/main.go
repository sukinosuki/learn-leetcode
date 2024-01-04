package main

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	if letters[n-1] <= target {
		return letters[0]
	}

	l := 0
	r := n - 1

	for l < r {

		mid := (r + l) / 2
		if letters[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return letters[r]
}

func main() {

}
