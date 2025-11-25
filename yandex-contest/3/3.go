package main

import (
	"bufio"
	"fmt"
	"os"
)

type DSU struct {
	p, r []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	r := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		r[i] = 0
	}
	return &DSU{p: p, r: r}
}

func (d *DSU) Find(x int) int {
	if d.p[x] != x {
		d.p[x] = d.Find(d.p[x])
	}
	return d.p[x]
}

func (d *DSU) Union(a, b int) {
	ra := d.Find(a)
	rb := d.Find(b)
	if ra == rb {
		return
	}
	if d.r[ra] < d.r[rb] {
		d.p[ra] = rb
	} else if d.r[ra] > d.r[rb] {
		d.p[rb] = ra
	} else {
		d.p[rb] = ra
		d.r[ra]++
	}
}

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	var n int
	_, err := fmt.Fscan(in, &n)
	if err != nil {
		return
	}

	queries := make([][]string, n)

	dsu := NewDSU(n)
	firstOccur := make(map[string]int)

	for i := 0; i < n; i++ {
		var m int
		fmt.Fscan(in, &m)
		queries[i] = make([]string, 0, m)
		for j := 0; j < m; j++ {
			var w string
			fmt.Fscan(in, &w)
			queries[i] = append(queries[i], w)
			if idx, ok := firstOccur[w]; ok {
				dsu.Union(idx, i)
			} else {
				firstOccur[w] = i
			}
		}
	}

	compWords := make(map[int]map[string]struct{})
	for i := 0; i < n; i++ {
		root := dsu.Find(i)
		if _, ok := compWords[root]; !ok {
			compWords[root] = make(map[string]struct{})
		}
		for _, w := range queries[i] {
			compWords[root][w] = struct{}{}
		}
	}

	numContexts := 0
	maxSize := 0
	for _, m := range compWords {
		if len(m) > 0 {
			numContexts++
			if len(m) > maxSize {
				maxSize = len(m)
			}
		}
	}

	if numContexts == 0 && n > 0 {
		numContexts = n
		maxSize = 0
	}

	fmt.Fprintln(out, numContexts, maxSize)
}
