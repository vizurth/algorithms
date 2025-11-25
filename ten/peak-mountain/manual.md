### 852. Peak Index in a Mountain Array

```go
func peakIndexInMountainArray(arr []int) int {
    l, r := 0, len(arr) - 1
    for l < r {
        m := (l + r)/2
        if arr[m] > arr[m + 1]{
            r = m
        } else {
            l = m + 1
        }
    }
    return l
}
```

#### Описание алгоритма

Левый бинпоиск и просто условие

#### О большое

Время: `O(log N)`
Память: `O(1)`
