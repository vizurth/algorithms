### 74. Search a 2d Matrix
```go
func searchMatrix(matrix [][]int, target int) bool {
    rows, cols := len(matrix), len(matrix[0])
    l, r := 0, rows*cols

    for l < r {
        m := (l + r) / 2
        val := matrix[m/cols][m%cols]
        if val >= target {
            r = m
        } else {
            l = m + 1
        }
    }

    if l == rows*cols {
        return false
    }
    return matrix[l/cols][l%cols] == target
}
```
### Описание алгоритма
