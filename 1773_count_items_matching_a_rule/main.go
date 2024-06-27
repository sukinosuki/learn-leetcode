package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {

	index := 0
	if ruleKey == "color" {
		index = 1
	}
	if ruleKey == "name" {
		index = 2
	}

	ans := 0
	for i := range items {
		if items[i][index] == ruleValue {
			ans++
		}
	}

	return ans
}

func main() {

}
