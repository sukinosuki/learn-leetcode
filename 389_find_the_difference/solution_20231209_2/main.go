package main

func findTheDifference(s string, t string) byte {

	res := 0
	for index := range s {
		res ^= int(s[index]) ^ int(t[index])
	}

	return byte(res ^ int(t[len(s)]))
}

func main() {

}
