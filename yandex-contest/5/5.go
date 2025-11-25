package main

import (
	"bufio"
	"fmt"
	"os"
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
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	readInt := func() int {
		sign := 1
		val := 0
		c, err := in.ReadByte()
		for err == nil && (c < '0' || c > '9') && c != '-' {
			c, err = in.ReadByte()
		}
		if err != nil {
			return 0
		}
		if c == '-' {
			sign = -1
			c, _ = in.ReadByte()
		}
		for err == nil && c >= '0' && c <= '9' {
			val = val*10 + int(c-'0')
			c, err = in.ReadByte()
		}
		return val * sign
	}

	n := readInt()
	k := readInt()

	graph = make([][]Edge, n+1)
	visited = make([]bool, n+1)
	tin = make([]int, n+1)
	low = make([]int, n+1)
	bridges = make(map[int]bool)
	timer = 0

	edges := make([][2]int, k)
	edgeExists := make(map[[2]int]bool)

	for i := 0; i < k; i++ {
		u := readInt()
		v := readInt()
		graph[u] = append(graph[u], Edge{v, i})
		graph[v] = append(graph[v], Edge{u, i})
		edges[i] = [2]int{u, v}
		a, b := u, v
		if a > b {
			a, b = b, a
		}
		edgeExists[[2]int{a, b}] = true
	}

	for i := 1; i <= n; i++ {
		if !visited[i] {
			dfs(i, -1)
		}
	}

	if len(bridges) == 0 {
		fmt.Fprintln(out, 0)
		return
	}

	comp := make([]int, n+1)
	visitedComp := make([]bool, n+1)
	compNum := 0
	compVertices := make([][]int, 1)

	for i := 1; i <= n; i++ {
		if !visitedComp[i] {
			compNum++
			compVertices = append(compVertices, []int{})
			bfsCollect(i, compNum, comp, visitedComp, &compVertices)
		}
	}

	compGraph := make([]map[int]bool, compNum+1)
	for i := 1; i <= compNum; i++ {
		compGraph[i] = make(map[int]bool)
	}
	for id := range bridges {
		u := edges[id][0]
		v := edges[id][1]
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
		fmt.Fprintln(out, 0)
		return
	}
	result := (leaves + 1) / 2
	fmt.Fprintln(out, result)

	preferRep := make([]int, compNum+1)
	for cid := 1; cid <= compNum; cid++ {
		preferRep[cid] = -1
	}
	for id := range bridges {
		u := edges[id][0]
		v := edges[id][1]
		cu := comp[u]
		cv := comp[v]
		if preferRep[cu] == -1 {
			preferRep[cu] = u
		}
		if preferRep[cv] == -1 {
			preferRep[cv] = v
		}
	}

	findPair := func(compA, compB int) (int, int) {
		if preferRep[compA] != -1 && preferRep[compB] != -1 {
			a, b := preferRep[compA], preferRep[compB]
			x, y := a, b
			if x > y {
				x, y = y, x
			}
			if !edgeExists[[2]int{x, y}] {
				return a, b
			}
		}
		A := compVertices[compA]
		B := compVertices[compB]
		if len(A) > len(B) {
			A, B = B, A
		}
		for _, va := range A {
			for _, vb := range B {
				x, y := va, vb
				if x > y {
					x, y = y, x
				}
				if !edgeExists[[2]int{x, y}] {
					return va, vb
				}
			}
		}
		return compVertices[compA][0], compVertices[compB][0]
	}

	for i := 0; i+1 < leaves; i += 2 {
		c1 := leafList[i]
		c2 := leafList[i+1]
		v1, v2 := findPair(c1, c2)
		fmt.Fprintf(out, "%d %d\n", v1, v2)
	}
	if leaves%2 == 1 {
		c1 := leafList[leaves-1]
		c2 := leafList[0]
		v1, v2 := findPair(c1, c2)
		fmt.Fprintf(out, "%d %d\n", v1, v2)
	}
}

func dfs(v, parentEdge int) {
	visited[v] = true
	tin[v] = timer
	low[v] = timer
	timer++

	for _, e := range graph[v] {
		to := e.to
		id := e.id
		if id == parentEdge {
			continue
		}
		if visited[to] {
			if tin[to] < low[v] {
				low[v] = tin[to]
			}
		} else {
			dfs(to, id)
			if low[to] < low[v] {
				low[v] = low[to]
			}
			if low[to] > tin[v] {
				bridges[id] = true
			}
		}
	}
}

func bfsCollect(start, compNum int, comp []int, visitedComp []bool, compVertices *[][]int) {
	queue := make([]int, 0, 16)
	queue = append(queue, start)
	visitedComp[start] = true
	comp[start] = compNum
	(*compVertices)[compNum] = append((*compVertices)[compNum], start)
	head := 0
	for head < len(queue) {
		v := queue[head]
		head++
		for _, e := range graph[v] {
			to := e.to
			id := e.id
			if !visitedComp[to] && !bridges[id] {
				visitedComp[to] = true
				comp[to] = compNum
				(*compVertices)[compNum] = append((*compVertices)[compNum], to)
				queue = append(queue, to)
			}
		}
	}
}
