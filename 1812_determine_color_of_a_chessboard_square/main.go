package main

import "fmt"

func squareIsWhite(coordinates string) bool {

	if (coordinates[0]-'a')%2 == 0 {

		return (coordinates[1]-'0')%2 == 0
	}

	return (coordinates[1]-'0')%2 == 1
}

func main() {
	// false
	//coordinate := "a1"
	// true
	coordinate := "h3"
	res := squareIsWhite(coordinate)

	fmt.Printf("res: %v\n", res)
}
