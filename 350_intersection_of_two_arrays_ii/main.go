package main

func intersect(nums1 []int, nums2 []int) []int {

	ans := make([]int, 0)
	m1 := make(map[int]int)

	for _, v := range nums1 {
		m1[v]++
	}

	for _, v := range nums2 {
		count, ok := m1[v]
		if ok && count > 0 {
			ans = append(ans, v)
			m1[v]--
		}

	}

	return ans
}

func main() {

}
