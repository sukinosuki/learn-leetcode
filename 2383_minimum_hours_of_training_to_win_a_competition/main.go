package main

import "fmt"

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {

	ans := 0
	for i := range energy {
		if energy[i] >= initialEnergy {
			diff := energy[i] - initialEnergy + 1
			ans += diff
			initialEnergy += diff
		}

		if experience[i] >= initialExperience {
			diff := experience[i] - initialExperience + 1
			ans += diff
			initialExperience += diff
		}

		initialEnergy -= energy[i]
		initialExperience += experience[i]
	}

	return ans
}

func main() {
	// 8
	initialEnergy := 5
	initialExperience := 3
	energy := []int{1, 4, 3, 2}
	experience := []int{2, 6, 3, 1}

	// 0
	//initialEnergy := 2
	//initialExperience := 4
	//energy := []int{1}
	//experience := []int{3}

	// energy: 4
	// exp:1+
	//initialEnergy := 1
	//initialExperience := 1
	//energy := []int{1, 1, 1, 1}
	//experience := []int{1, 1, 1, 50}

	// 0
	//initialEnergy := 100
	//initialExperience := 100
	//energy := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//experience := []int{1, 2, 3, 1, 2, 3, 1, 2, 10}
	res := minNumberOfHours(initialEnergy, initialExperience, energy, experience)
	fmt.Printf("res: %v\n", res)
}
