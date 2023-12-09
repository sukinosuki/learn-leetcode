package main

func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 0; i <= n; i++ {
		subBinary := i >> 1

		ans[i] = ans[subBinary] + i&1
	}

	return ans
}

func main() {

}
