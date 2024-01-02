package main

func judgeCircle(moves string) bool {

	//m:= make(map[int]int)
	countTop := 0
	countRight := 0

	for _, v := range moves {
		if v == 'U' {
			countTop++
		} else if v == 'R' {
			countRight++
		} else if v == 'L' {
			countRight--

		} else if v == 'D' {
			countTop--
		}
	}

	return countTop == 0 && countRight == 0
}

func main() {

}
