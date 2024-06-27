package main

func findCenter(edges [][]int) int {

	if edges[1][0] == edges[0][0] || edges[1][0] == edges[0][1] {

		return edges[1][0]
	}

	return edges[1][1]
}

func main() {

}
