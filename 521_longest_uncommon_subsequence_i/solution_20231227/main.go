package main

func findLUSlength(a string, b string) int {
	n1 := len(a)
	n2 := len(b)

	if a == b {
		return -1
	}

	if n1 > n2 {
		return n1
	}

	return n2
}

func main() {

}
