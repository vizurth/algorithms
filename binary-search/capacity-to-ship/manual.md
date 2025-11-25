### 1011. Capacity To Ship Packages Withn D Days
```go
func check(m, days int, weights []int) bool {
    deep := 0
    count := 0
    for _, elem := range weights {
        if deep + elem > m {
            count++
            deep = 0
        }
        deep += elem
    }
    count++
    return count <= days
}


func shipWithinDays(weights []int, days int) int {
    l, r  := 0,0

    for _ , val := range weights {
        r += val
        l = max(l , val)
    }

    for l < r {
        m := (l + r) / 2
        if check(m, days, weights) {
            r = m
        } else {
            l = m + 1
        }
    }
    return l
}
```
### Описание алгоритма
