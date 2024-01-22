package main

import (
	"fmt"
)

func isLongPressedName(name string, typed string) bool {

	n := len(name)
	m := len(typed)
	if n > m {
		return false
	}

	l1 := 0
	l2 := 0
	for l1 < n && l2 < m {
		if name[l1] == typed[l2] {
			l1++
			l2++
			continue
		}
		if l1 == 0 {
			return false
		}
		if l2 > 0 && typed[l2] != typed[l2-1] {
			return false
		}
		l2++
	}

	for l2 < m {
		if typed[l2] != typed[l2-1] {
			return false
		}
		l2++
	}

	return l1 == n
}

func main() {
	//name := "alex"
	//typed := "aaleex"

	//name := "alex"
	//typed := "aaleexa"

	//name := "vtkgn"
	//typed := "vttkgnn"

	name := "zlexya"
	typed := "aazlllllllllllllleexxxxxxxxxxxxxxxya"

	//name := "saeed"
	//typed := "ssaaedd"
	//typed := "ssaaeddedd"
	res := isLongPressedName(name, typed)

	fmt.Printf("res: %v\n", res)
}
