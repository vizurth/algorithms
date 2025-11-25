### 56. Merge Intervals
```go
func isOverlapping(a, b []int) bool {
    return max(a[0], b[0]) <= min(a[1], b[1])
}

func mergeTwo(a, b []int) []int {
    return []int{a[0], max(a[1], b[1])}
}

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func (a, b int) bool {
        return intervals[a][0] < intervals[b][0]
    })

    result := make([][]int, 1, len(intervals))
    result[0] = intervals[0]
    for i := 1; i < len(intervals); i++ {
        if isOverlapping(result[len(result) - 1], intervals[i]) {
            result[len(result) - 1] = mergeTwo(result[i-1], intervals[i]) 
        } else {
            result = append(result, intervals[i])
        }
    }  

    return result
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
Алгоритм использует классический жадный подход:

Сначала интервалы сортируются по началу.

Затем мы проходим по ним слева направо и объединяем те, которые пересекаются.

Если интервал не пересекается с предыдущим — добавляем его как новый.
