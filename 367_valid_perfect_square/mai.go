package main

func isPerfectSquare(num int) bool {

	x := 1
	square := 1
	for square <= num {
		if square == num {
			return true
		}
		x += 1

		square = x * x
	}

	return false
}

func main() {

}
