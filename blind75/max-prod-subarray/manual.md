### 20. Maximum Product Subarray
```go
func maxProduct(nums []int) int {
    res := nums[0]
    currMax, currMin := nums[0], nums[0]

    for i := 1; i < len(nums); i++ {
        val := nums[i]
        if val < 0 {
            currMax, currMin = currMin, currMax
        }
        currMax = max(val, val*currMax)
        currMin = min(val, val*currMin)
        res = max(res, currMax)
    }

    return res
}
```
### Описание алгоритма
