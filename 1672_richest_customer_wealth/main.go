package main

func maximumWealth(accounts [][]int) int {

	ans := 0

	for i := range accounts {

		sum := 0
		for j := range accounts[i] {
			sum += accounts[i][j]
		}
		if sum > ans {
			ans = sum
		}
	}

	return ans
}

func main() {

}
