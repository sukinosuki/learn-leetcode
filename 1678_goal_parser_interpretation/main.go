package main

import "fmt"

func interpret(command string) string {

	var ans []byte

	i := 0
	n := len(command)

	for i < n {

		if command[i] == 'G' {
			ans = append(ans, command[i])
			i++
			continue
		}
		if command[i:i+2] == "()" {
			ans = append(ans, 'o')
			i += 2
			continue
		}

		ans = append(ans, "al"...)
		i += 4
	}

	return string(ans)
}

func main() {

	//  Goal
	//command := "G()(al)"

	// Gooooal
	//command := "G()()()()(al)"
	command := "(al)G(al)()()G"
	res := interpret(command)

	fmt.Printf("res: %v\n", res)
}
