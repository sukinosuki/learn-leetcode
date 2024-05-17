package main

func defangIPaddr(address string) string {
	var ans []byte

	for i := range address {
		if address[i] == '.' {
			ans = append(ans, '[')
			ans = append(ans, '.')
			ans = append(ans, ']')
			continue
		}

		ans = append(ans, address[i])
	}

	return string(ans)
}

func main() {

}
