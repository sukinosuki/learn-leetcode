package main

func canBeEqual(target []int, arr []int) bool {

	n := len(arr)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[target[i]]++
		m[arr[i]]--
	}

	for i := range m {
		if m[i] < 0 {
			return false
		}
	}

	return true
}
func main() {

}
