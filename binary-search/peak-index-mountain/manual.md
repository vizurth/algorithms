### 852. Peak Index in a Mountain Array
```go
func leftBinSearch(nums []int) int {
    l, r := 0, len(nums) - 1
    for l < r {
        m := (l + r) / 2
        if nums[m] > nums[m+1] {
            r = m
        } else {
            l = m + 1
        }
    }
    return l
}

func peakIndexInMountainArray(arr []int) int {
    return leftBinSearch(arr)
}
```
### Описание алгоритма
