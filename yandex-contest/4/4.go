package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	to, id int
}

var (
	graph   [][]Edge
	visited []bool
	tin     []int
	low     []int
	timer   int
	bridges map[int]bool
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])

	graph = make([][]Edge, n+1)
	visited = make([]bool, n+1)
	tin = make([]int, n+1)
	low = make([]int, n+1)
	bridges = make(map[int]bool)
	timer = 0

	edges := make([][2]int, k)
	for i := 0; i < k; i++ {
		line, _ = reader.ReadString('\n')
		parts = strings.Fields(line)
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		graph[u] = append(graph[u], Edge{v, i})
		graph[v] = append(graph[v], Edge{u, i})
		edges[i] = [2]int{u, v}
	}

	for i := 1; i <= n; i++ {
		if !visited[i] {
			dfsIter(i)
		}
	}

	if len(bridges) == 0 {
		writer.WriteString("0\n")
		return
	}

	comp := make([]int, n+1)
	visitedComp := make([]bool, n+1)
	compNum := 0
	for i := 1; i <= n; i++ {
		if !visitedComp[i] {
			compNum++
			bfs(i, compNum, comp, visitedComp)
		}
	}

	compGraph := make([]map[int]bool, compNum+1)
	for i := 1; i <= compNum; i++ {
		compGraph[i] = make(map[int]bool)
	}
	for edgeID := range bridges {
		u := edges[edgeID][0]
		v := edges[edgeID][1]
		cu := comp[u]
		cv := comp[v]
		if cu != cv {
			compGraph[cu][cv] = true
			compGraph[cv][cu] = true
		}
	}

	leafList := []int{}
	for i := 1; i <= compNum; i++ {
		if len(compGraph[i]) == 1 {
			leafList = append(leafList, i)
		}
	}

	leaves := len(leafList)
	if leaves == 0 {
		writer.WriteString("0\n")
		return
	}

	result := (leaves + 1) / 2
	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')

	for i := 0; i < leaves/2; i++ {
		v1 := findVertexInComp(comp, leafList[2*i], n)
		v2 := findVertexInComp(comp, leafList[2*i+1], n)
		writer.WriteString(strconv.Itoa(v1))
		writer.WriteByte(' ')
		writer.WriteString(strconv.Itoa(v2))
		writer.WriteByte('\n')
	}
	if leaves%2 == 1 {
		v1 := findVertexInComp(comp, leafList[leaves-1], n)
		v2 := findVertexInComp(comp, leafList[0], n)
		writer.WriteString(strconv.Itoa(v1))
		writer.WriteByte(' ')
		writer.WriteString(strconv.Itoa(v2))
		writer.WriteByte('\n')
	}
}

func dfsIter(start int) {
	type Frame struct {
		v          int
		parentEdge int
		i          int 
	}
	stack := []Frame{{start, -1, 0}}

	for len(stack) > 0 {
		top := &stack[len(stack)-1]
		v := top.v

		if !visited[v] {
			visited[v] = true
			tin[v] = timer
			low[v] = timer
			timer++
		}

		if top.i < len(graph[v]) {
			e := graph[v][top.i]
			top.i++

			to := e.to
			id := e.id
			if id == top.parentEdge {
				continue
			}

			if visited[to] {
				if tin[to] < tin[v] {
					if low[v] > tin[to] {
						low[v] = tin[to]
					}
				}
			} else {
				stack = append(stack, Frame{to, id, 0})
			}
			continue
		}

		stack = stack[:len(stack)-1]
		if len(stack) > 0 {
			parent := &stack[len(stack)-1]
			if low[parent.v] > low[v] {
				low[parent.v] = low[v]
			}
			if low[v] > tin[parent.v] {
				bridges[top.parentEdge] = true
			}
		}
	}
}


func dfsIter(start int) {
	type frame struct {
		v          int
		parentEdge int
		nextIndex  int
	}
	stack := []frame{}
	stack = append(stack, frame{start, -1, 0})

	for len(stack) > 0 {
		f := &stack[len(stack)-1]
		v := f.v

		if !visited[v] {
			visited[v] = true
			tin[v] = timer
			low[v] = timer
			timer++
		}

		if f.nextIndex < len(graph[v]) {
			e := graph[v][f.nextIndex]
			f.nextIndex++
			to := e.to
			id := e.id
			if id == f.parentEdge {
				continue
			}
			if visited[to] {
				if low[v] > tin[to] {
					low[v] = tin[to]
				}
			} else {
				stack = append(stack, frame{to, id, 0})
			}
		} else {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				parent := &stack[len(stack)-1]
				if low[parent.v] > low[popped.v] {
					low[parent.v] = low[popped.v]
				}
				if low[popped.v] > tin[parent.v] {
					bridges[popped.parentEdge] = true
				}
			}
		}
	}
}

func bfs(start, compNum int, comp []int, visitedComp []bool) {
	queue := []int{start}
	visitedComp[start] = true
	comp[start] = compNum

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, edge := range graph[v] {
			to := edge.to
			id := edge.id
			if !visitedComp[to] && !bridges[id] {
				visitedComp[to] = true
				comp[to] = compNum
				queue = append(queue, to)
			}
		}
	}
}

func findVertexInComp(comp []int, compNum, n int) int {
	for i := 1; i <= n; i++ {
		if comp[i] == compNum {
			return i
		}
	}
	return 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
