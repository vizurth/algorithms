### 1283. Find the Smallest Divisor Given a Threshold
```go
func check(div, threshold int, nums []int) bool {
    sum := 0
    for _, num := range nums {
        sum += (num + div - 1) / div
    }
    return sum <= threshold
}

func smallestDivisor(nums []int, threshold int) int {
    l, r := 1, 0
    for _, num := range nums {
        if num > r {
            r = num
        }
    }

    for l < r {
        m := (l + r) / 2
        if check(m, threshold, nums) {
            r = m
        } else {
            l = m + 1
        }
    }

    return l
}
```
### Описание алгоритма
