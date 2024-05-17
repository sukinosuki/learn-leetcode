package main

func findNumbers(nums []int) int {

	ans := 0
	for _, num := range nums {
		if num >= 10 && num < 100 {
			ans++
			continue
		}
		if num >= 1000 && num < 10000 {
			ans++
			continue
		}
		if num >= 100000 && num < 1000000 {
			ans++
		}
	}

	return ans
}

func main() {

}
