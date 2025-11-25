### 200. Number of Islands
```go
func dfs(grid [][]byte, i, j int) {
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0'{
        return
    }
    grid[i][j] = '0'
    dfs(grid, i + 1, j)
    dfs(grid, i - 1, j)
    dfs(grid, i, j + 1)
    dfs(grid, i, j - 1)
}

func bfs(grid [][]byte, i, j int) {
    queue := [][]int{}
    queue = append(queue, []int{i, j})

    direction := [][]int {
        {1,0}, {-1,0}, {0,1}, {0,-1},
    }

    grid[i][j] = '0' // отметили посещенную

    for len(queue) > 0 {
        item := queue[0]
        queue = queue[1:]

        for _, d := range direction {
            i1, j1 := d[0] + item[0], d[1]+item[1]

            if i1 >= 0 && i1 < len(grid) && j1 >= 0 && j1 < len(grid[0]) && grid[i1][j1] == '1' {
                grid[i1][j1] = '0'
                queue = append(queue, []int{i1, j1})
            }
        }
    }
}

func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    count := 0

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                count++
                bfs(grid, i, j)
                //dfs(grid, i, j)
            }
        }
    }


    return count
}
```
### Описание алгоритма
У нас есть матрица из 0 и 1.
Каждая группа связанных по вертикали или горизонтали единиц — это остров.
Нужно просто посчитать, сколько таких групп есть.
То есть:
Каждый раз, когда ты находишь клетку с '1', это новый остров.
Потом ты «заливаешь» (обходишь) весь этот остров, превращая все связанные '1' в '0', чтобы не считать их снова.
Продолжаешь, пока не пройдёшь всю сетку.