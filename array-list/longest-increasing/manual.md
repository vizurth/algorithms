### 674. Longest Continuos Increasing Subsequence
```go
func findLengthOfLCIS(nums []int) int {
    count := 0
    l := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1]{
            count = max(count, i-l)
        } else {
            l = i
        }
    }

    return count + 1
}
```
### Описание алгоритма:
Будем смотреть на предыдущий и считать максимум иначе сдвигаем указадель `(l = i)`считаем длину подпоследоваленость == `i - l`