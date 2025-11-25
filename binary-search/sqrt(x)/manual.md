### 69. Sqrt(x)
```go
func mySqrt(x int) int {
    if x < 2 {
        return x
    }
    l, r := 0, x
    for l < r {
        m := (l + r) / 2
        if int64(m)*int64(m) > int64(x) {
            r = m
        } else {
            l = m + 1
        }
    }
    return l - 1
}
```
### Описание алгоритма:
