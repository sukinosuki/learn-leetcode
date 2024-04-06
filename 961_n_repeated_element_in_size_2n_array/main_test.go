package main

import "testing"

func TestRepeated1(t *testing.T) {
	nums := []int{1, 2, 3, 3}
	res := RepeatedNTimes(nums)
	if res != 3 {
		t.Errorf("result was incorrect, got: %d, want: %d.", res, 3)
	}
}
