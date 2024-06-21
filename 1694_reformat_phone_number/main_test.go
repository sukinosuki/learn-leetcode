package main

import (
	"testing"
)

func TestReformatNumber(t *testing.T) {
	number := "1-23-45 6"

	res := ReformatNumber(number)

	expect := "123-456"
	if res != expect {
		t.Errorf("失败: 期望: %s, 结果: %s", expect, res)
	}
	//
	//fmt.Printf("通过")
}
