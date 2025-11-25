package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	line, _ = reader.ReadString('\n')
	valueParts := strings.Fields(line)
	values := make([]int, n+1)
	for i := 1; i <= n; i++ {
		values[i], _ = strconv.Atoi(valueParts[i-1])
	}

	graph := make([][]int, n+1)
	for i := 0; i < m; i++ {
		line, _ = reader.ReadString('\n')
		edge := strings.Fields(line)
		a, _ := strconv.Atoi(edge[0])
		b, _ := strconv.Atoi(edge[1])
		graph[a] = append(graph[a], b)
	}

	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	
	var dfs func(int) int
	dfs = func(node int) int {
		if dp[node] != -1 {
			return dp[node]
		}
		
		dp[node] = values[node]
		
		maxFromNeighbors := 0
		for _, next := range graph[node] {
			maxFromNeighbors = max(maxFromNeighbors, dfs(next))
		}
		
		dp[node] += maxFromNeighbors
		return dp[node]
	}
	
	result := dfs(1)
	
	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}