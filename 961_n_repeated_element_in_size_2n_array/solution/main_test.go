package main

import "testing"

func TestRepeatedNTimes(t *testing.T) {
	nums := []int{1, 2, 3, 3} // 3

	times := repeatedNTimes(nums)

	if times != 3 {
		t.Errorf("result was incorrect, want: %d, get: %d", 3, times)
	}
}
