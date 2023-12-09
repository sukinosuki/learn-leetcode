package main

func findTheDifference(s string, t string) byte {
	sum := 0
	for _, char := range s {
		sum -= int(char)
	}

	for _, char := range t {
		sum += int(char)
	}

	return byte(sum)
}

func main() {

}
