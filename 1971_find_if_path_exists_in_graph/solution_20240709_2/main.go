package main

import "fmt"

func validPath(n int, edges [][]int, source int, destination int) bool {

	var adj = make([][]int, n)

	for i := range edges {
		v1, v2 := edges[i][0], edges[i][1]
		adj[v1] = append(adj[v1], v2)
		adj[v2] = append(adj[v2], v1)
	}

	visited := make([]bool, n)

	return dfs(source, destination, adj, visited)
}

func dfs(source int, destination int, adj [][]int, visited []bool) bool {
	if source == destination {
		return true
	}
	arr := adj[source]
	visited[source] = true

	for _, num := range arr {
		if !visited[num] && dfs(num, destination, adj, visited) {
			return true
		}
	}
	return false
}

func main() {
	n := 3
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 0},
	}
	source := 0
	destination := 2
	res := validPath(n, edges, source, destination)

	fmt.Printf("res: %v\n", res)
}
