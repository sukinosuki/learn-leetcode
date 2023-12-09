package main

func firstBadVersion(n int) int {

	l := 1
	r := n

	// 1 2 3 4 5 6 7
	// t t f f f f f
	for l < r {
		mid := l + (r-l)/2
		ok := isBadVersion(mid)
		// 是错误版本
		if ok {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func isBadVersion(version int) bool {

	return version >= 2
}

func main() {
	firstBadVersion(3)
}
