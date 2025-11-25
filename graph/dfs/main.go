package main

import "fmt"


func Bfs(startNode int, visited *map[int]bool, graph map[int][]int) {
	queue := []int{startNode}
	(*visited)[startNode] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[node] {
			if !(*visited)[neighbor] {
				(*visited)[neighbor] = true
				queue = append(queue, neighbor)
				// какая-то логика

			}
		}
	}
}

func Dfs(graph map[string][]string, node string, visited map[string]bool) {
	if visited[node] {
		return
	}

	visited[node] = true
	fmt.Println(node)

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			Dfs(graph, neighbor, visited)
		}
	}
}