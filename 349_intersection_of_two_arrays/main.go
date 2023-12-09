package main

func intersection(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)

	m1 := make(map[int]bool)
	for _, v := range nums1 {
		m1[v] = true
	}

	for _, v := range nums2 {
		if flag, ok := m1[v]; ok && flag {
			ans = append(ans, v)
			m1[v] = false
		}
	}

	return ans
}

func main() {

}
