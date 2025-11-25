### Subarray Sum Equals K
```go
func subarraySum(nums []int, k int) int {
    res, summ := 0, 0
    m := make(map[int]int)
    m[0] = 1
    for _, v := range nums {
        summ += v
        res += m[summ - k]
        m[summ]++
    }
    return res
}
```
### Описание алгоритма
