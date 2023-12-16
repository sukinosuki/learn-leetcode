package main

func maxProfit(prices []int) int {

	minPrice := prices[0]
	_maxProfit := 0

	for i := 0; i < len(prices); i++ {
		_maxProfit = max(_maxProfit, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}

	return _maxProfit
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {

}
