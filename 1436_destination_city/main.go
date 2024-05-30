package main

import "fmt"

func destCity(paths [][]string) string {

	m := make(map[string]string)

	for i := range paths {
		m[paths[i][0]] = paths[i][1]
	}

	key := paths[0][1]

	for true {
		value, ok := m[key]
		if ok {
			key = value
			continue
		}
		break
	}

	return key
}

func main() {
	// Sao Paulo
	//paths := [][]string{
	//	{"London", "New York"},
	//	{"New York", "Lima"},
	//	{"Lima", "Sao Paulo"},
	//}

	// wxAscRuzOl
	//paths := [][]string{
	//	{"pYyNGfBYbm", "wxAscRuzOl"},
	//	{"kzwEQHfwce", "pYyNGfBYbm"},
	//}

	// OAxMijOZgW
	paths := [][]string{
		{"qMTSlfgZlC", "ePvzZaqLXj"},
		{"xKhZXfuBeC", "TtnllZpKKg"},
		{"ePvzZaqLXj", "sxrvXFcqgG"},
		{"sxrvXFcqgG", "xKhZXfuBeC"},
		{"TtnllZpKKg", "OAxMijOZgW"},
	}

	res := destCity(paths)

	fmt.Printf("res: %v\n", res)
}
