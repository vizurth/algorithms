### 34. Find First and Last Position of Element in Sorted Array
```go
func leftSearch(l, r, target int, nums []int) int {
    for l < r {
        m := (l + r) / 2
        if nums[m] >= target {
            r = m
        } else {
            l = m + 1
        }
    }
    return l
}

func rightSearch(l, r, target int, nums []int) int {
    for l < r {
        m := (l + r + 1) / 2
        if nums[m] <= target {
            l = m
        } else {
            r = m - 1
        }
    }
    return l
}

func searchRange(nums []int, target int) []int {
    if len(nums) == 0{
        return []int{-1,-1}
    }
    l, r := leftSearch(0, len(nums)-1, target, nums), rightSearch(0, len(nums)-1, target, nums)
    if nums[l] == target && nums[r] == target{
        return []int{l,r}
    }
    return []int{-1,-1}
}
```
### Описание алгоритма