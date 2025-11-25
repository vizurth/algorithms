package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Cell struct {
	r, c int
	t    int
}

type PriorityQueue []Cell

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].t < pq[j].t }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Cell))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var n, m int
	if _, err := fmt.Fscan(in, &n, &m); err != nil {
		return
	}

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &grid[i][j])
		}
	}

	const INF = int(1<<60)
	flood := make([][]int, n)
	for i := 0; i < n; i++ {
		flood[i] = make([]int, m)
		for j := 0; j < m; j++ {
			flood[i][j] = -1
		}
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				flood[i][j] = 0
				heap.Push(pq, Cell{i, j, 0})
			}
		}
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Cell)
		if flood[cur.r][cur.c] != -1 && cur.t != flood[cur.r][cur.c] {
			continue
		}
		for _, d := range dirs {
			ni, nj := cur.r+d[0], cur.c+d[1]
			if ni < 0 || ni >= n || nj < 0 || nj >= m {
				continue
			}
			nt := max(cur.t, grid[ni][nj])
			if flood[ni][nj] == -1 || nt < flood[ni][nj] {
				flood[ni][nj] = nt
				heap.Push(pq, Cell{ni, nj, nt})
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j > 0 {
				fmt.Fprint(out, " ")
			}
			if flood[i][j] == -1 {
				fmt.Fprint(out, grid[i][j])
			} else {
				fmt.Fprint(out, flood[i][j])
			}
		}
		fmt.Fprintln(out)
	}
}
