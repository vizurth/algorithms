### 452. Minimum Number of Arrows to Burst Balloons
```go
func isOverlapping(a, b []int) bool {
    return max(a[0], b[0]) <= min(a[1], b[1])
}

func overlap(a, b []int) []int {
    return []int{max(a[0], b[0]), min(a[1], b[1])}
}

func findMinArrowShots(intervals [][]int) int {
    sort.Slice(intervals, func (a, b int) bool {
        return intervals[a][0] < intervals[b][0]
    })

    result := make([][]int, 1, len(intervals))
    result[0] = intervals[0]
    for i := 1; i < len(intervals); i++ {
        if isOverlapping(result[len(result) - 1], intervals[i]) {
            result[len(result) - 1] = mergeTwo(result[len(result) - 1], intervals[i]) 
        } else {
            result = append(result, intervals[i])
        }
    }  

    return len(result)
}


func max(a, b int) int{
    if a > b {
        return a
    } 
    return b
}

func min(a, b int) int{
    if a > b {
        return b
    } 
    return a
}
```
### Описание алгоритма
Если два шара пересекаются (их интервалы накладываются), то одна стрела может лопнуть оба.
Если интервалы не пересекаются, нужна новая стрела.

Таким образом, задача сводится к подсчёту количества непересекающихся групп интервалов.
